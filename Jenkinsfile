node ('infrastructure'){
  checkout scm

  try {
    stage 'test'
    sh 'make test'

    stage 'build'
    sh 'docker build --tag ${REGISTRY_HOST}:${REGISTRY_PORT}/go-ip-microservice --tag ${REGISTRY_HOST}:${REGISTRY_PORT}/go-ip-microservice:latest --tag ${REGISTRY_HOST}:${REGISTRY_PORT}/go-ip-microservice:$(date "+%d-%m-%Y_%H-%M-%S") .'

    stage 'login to registry'
    sh 'docker login --username ${REGISTRY_USERNAME} --password ${REGISTRY_PASSWORD} ${REGISTRY_HOST}:${REGISTRY_PORT}'

    stage 'push'
    sh 'docker push ${REGISTRY_HOST}:${REGISTRY_PORT}/go-ip-microservice'

    stage 'deploy'
    sh 'ansible-playbook /infrastructure/ansible/role.yml -i /infrastructure/ansible/hosts/${ENV_ENVIRONMENT} -e "HOST=${DEPLOYMENT_HOST}" -e "ROLE=$(pwd)/ansible/roles/deploy"'

    stage 'verify deployment'
    sh 'ansible-playbook /infrastructure/ansible/role.yml -i /infrastructure/ansible/hosts/${ENV_ENVIRONMENT} -e "HOST=${DEPLOYMENT_HOST}" -e "ROLE=$(pwd)/ansible/roles/deploy/tests"'
  } finally {
    stage 'teardown'
    sh 'docker-compose stop'
    sh 'docker-compose rm --all --force -v'
  }
}