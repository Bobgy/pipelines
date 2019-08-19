#!/bin/bash
#
# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex

usage()
{
    echo "usage: deploy.sh
    [--platform             the deployment platform. Valid values are: [gcp, minikube]. Default is gcp.]
    [--project              the gcp project. Default is ml-pipeline-test. Only used when platform is gcp.]
    [--workflow_file        the file name of the argo workflow to run]
    [--test_result_bucket   the gcs bucket that argo workflow store the result to. Default is ml-pipeline-test
    [--test_result_folder   the gcs folder that argo workflow store the result to. Always a relative directory to gs://<gs_bucket>/[PULL_SHA]]
    [--timeout              timeout of the tests in seconds. Default is 1800 seconds. ]
    [-h help]"
}

PLATFORM=gcp
PROJECT=ml-pipeline-test
TEST_RESULT_BUCKET=ml-pipeline-test
TIMEOUT_SECONDS=6000
NAMESPACE=kubeflow

while [ "$1" != "" ]; do
    case $1 in
             --platform )             shift
                                      PLATFORM=$1
                                      ;;
             --project )              shift
                                      PROJECT=$1
                                      ;;
             --workflow_file )        shift
                                      WORKFLOW_FILE=$1
                                      ;;
             --test_result_bucket )   shift
                                      TEST_RESULT_BUCKET=$1
                                      ;;
             --test_result_folder )   shift
                                      TEST_RESULT_FOLDER=$1
                                      ;;
             --timeout )              shift
                                      TIMEOUT_SECONDS=$1
                                      ;;
             -h | --help )            usage
                                      exit
                                      ;;
             * )                      usage
                                      exit 1
    esac
    shift
done

# Choose gcloud CLI project
gcloud config set project ${PROJECT}

# Variables
PROJECT_NUMBER=$(gcloud projects describe ${PROJECT} --format='value(projectNumber)')
# Default service account
# ref: https://cloud.google.com/compute/docs/access/service-accounts#compute_engine_default_service_account
VM_SERVICE_ACCOUNT="${PROJECT_NUMBER}-compute@developer.gserviceaccount.com"
GCR_IMAGE_BASE_DIR=gcr.io/${PROJECT}/${PULL_PULL_SHA}
TEST_RESULTS_GCS_DIR=gs://${TEST_RESULT_BUCKET}/${PULL_PULL_SHA}/${TEST_RESULT_FOLDER}
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

echo "presubmit test starts"
source "${DIR}/test-prep.sh"

# Deploy Kubeflow Pipelines lightweight deployment
source "${DIR}/deploy-cluster.sh"

# Install Argo CLI and test-runner service account
source "${DIR}/install-argo.sh"

IMAGE_BUILDER_ARG=""
# When project is not ml-pipeline-test, VMs need permission to fetch some images in gcr.io/ml-pipeline-test.
if [ "$PROJECT" != "ml-pipeline-test" ]; then
  COPIED_IMAGE_BUILDER_IMAGE=${GCR_IMAGE_BASE_DIR}/image-builder
  echo "Copy image builder image to ${COPIED_IMAGE_BUILDER_IMAGE}"
  yes | gcloud container images add-tag \
    gcr.io/ml-pipeline-test/image-builder:v20181128-0.1.3-rc.1-109-ga5a14dc-e3b0c4 \
    ${COPIED_IMAGE_BUILDER_IMAGE}:latest
  IMAGE_BUILDER_ARG="-p image-builder-image=${COPIED_IMAGE_BUILDER_IMAGE}"
fi

BUILT_IMAGES=$(gcloud container images list --repository=${GCR_IMAGE_BASE_DIR})
if
  test -n "$CACHE_IMAGES" && \
  echo "$BUILT_IMAGES" | grep  api-server && \
  echo "$BUILT_IMAGES" | grep frontend && \
  echo "$BUILT_IMAGES" | grep scheduledworkflow && \
  echo "$BUILT_IMAGES" | grep persistenceagent;
then
  echo "docker images for api-server, frontend, scheduledworkflow and \
    persistenceagent are already built in ${GCR_IMAGE_BASE_DIR}."
else
  echo "submitting argo workflow to build docker images for commit ${PULL_PULL_SHA}..."
  # Build Images
  ARGO_WORKFLOW=`argo submit ${DIR}/build_image.yaml \
  -p image-build-context-gcs-uri="$remote_code_archive_uri" \
  ${IMAGE_BUILDER_ARG} \
  -p api-image="${GCR_IMAGE_BASE_DIR}/api-server" \
  -p frontend-image="${GCR_IMAGE_BASE_DIR}/frontend" \
  -p scheduledworkflow-image="${GCR_IMAGE_BASE_DIR}/scheduledworkflow" \
  -p persistenceagent-image="${GCR_IMAGE_BASE_DIR}/persistenceagent" \
  -n ${NAMESPACE} \
  --serviceaccount test-runner \
  -o name
  `
  echo "build docker images workflow submitted successfully"
  source "${DIR}/check-argo-status.sh"
  echo "build docker images workflow completed"
fi

# Deploy the pipeline
source ${DIR}/deploy-pipeline-light.sh

echo "submitting argo workflow to run tests for commit ${PULL_PULL_SHA}..."
ARGO_WORKFLOW=`argo submit ${DIR}/${WORKFLOW_FILE} \
-p image-build-context-gcs-uri="$remote_code_archive_uri" \
${IMAGE_BUILDER_ARG} \
-p target-image-prefix="${GCR_IMAGE_BASE_DIR}/" \
-p test-results-gcs-dir="${TEST_RESULTS_GCS_DIR}" \
-p cluster-type="${CLUSTER_TYPE}" \
-n ${NAMESPACE} \
--serviceaccount test-runner \
-o name
`
echo "test workflow submitted successfully"
source "${DIR}/check-argo-status.sh"
echo "test workflow completed"
