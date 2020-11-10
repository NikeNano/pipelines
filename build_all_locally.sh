export API_SERVER_TA="api_server"
export CAHCE_SERVER_TAG="cache_server" 
export CAHCE_SERVER_TAG="cache_server"
export PERSISTENCEAGENT_SERVER_TAG="persistenceagent_server"
export SCHEDULEDWORKFLOW_SERVER_TAG="scheduledworkflow_server"
export VIEWERCONTROLLER_SERVER_TAG="viewerconroller_server"
export VISUALIZATION_SERVER_TAG="visualization_server"
export DOCKER_USERNAME='nikenano'


docker build \
      -t bazel:0.24.0 \
      -f Backend/Dockerfile.bazel .

docker build \
    -t "${DOCKER_USERNAME}"/"${API_SERVER_TAG}" \
    -f Backend/Dockerfile \
    . \
    --build-arg BAZEL_IMAGE=bazel:0.24.0


docker build \
    -t "${DOCKER_USERNAME}"/"${CAHCE_SERVER_TAG}" \
    -f Backend/Dockerfile.cacheserver  \
    . \
    --build-arg BAZEL_IMAGE=bazel:0.24.0

docker build \
    -t "${DOCKER_USERNAME}"/"${PERSISTENCEAGENT_SERVER_TAG}" \
    -f Backend/Dockerfile.persistenceagent \
    . \
    --build-arg BAZEL_IMAGE=bazel:0.24.0

docker build \
    -t "${DOCKER_USERNAME}"/"${SCHEDULEDWORKFLOW_SERVER_TAG}" \
    -f Backend/Dockerfile.scheduledworkflow\
    . \
    --build-arg BAZEL_IMAGE=bazel:0.24.0

docker build \
    -t "${DOCKER_USERNAME}"/"${VIEWERCONTROLLER_SERVER_TAG}" \
    -f Backend/Dockerfile.viewercontroller\
    . \
    --build-arg BAZEL_IMAGE=bazel:0.24.0

docker build \
    -t "${DOCKER_USERNAME}"/"${VISUALIZATION_SERVER_TAG}" \
    -f Backend/Dockerfile.visualization\
    . \
    --build-arg BAZEL_IMAGE=bazel:0.24.0
