# -*- mode: Python -*-
docker_build(
    'snapshotdb-image',
    '.',
    dockerfile='db/deployments/local/Dockerfile'
)

docker_build('api-image', '.', dockerfile='api/deployments/local/Dockerfile')

k8s_yaml([
 './db/deployments/local/deployment.yaml',
 './api/deployments/local/kubernetes.yaml',
])

k8s_resource(workload='snapshotdb', port_forwards=5432)
k8s_resource(workload='api', resource_deps=['snapshotdb'], port_forwards=8080)