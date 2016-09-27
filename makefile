SHELL=sh -x

test:
	docker-compose run test-http-interface

publish_nodejs_sdk:
	docker run \
		-v $$(pwd)/swagger.yaml:/swagger.yaml \
		-e NPM_PACKAGE_NAME=$${NPM_PACKAGE_NAME} \
		-e CODEGEN_SWAGGER_FILE=/swagger.yaml \
		-e CODEGEN_LANGUAGE=typescript-fetch \
		-e NPM_AUTH_TOKEN=$${NPM_AUTH_TOKEN} \
		21stio/swagger-codegen-npm-publish

push_go_sdk:
	docker run \
  	-v $$(pwd)/swagger.yaml:/swagger.yaml \
  	-v $${GITHUB_SSH_KEY}:/github \
  	-e SSH_KEY_FILE=/github \
  	-e GIT_HOST=github.com \
  	-e GIT_USERNAME=$${GIT_USERNAME} \
  	-e GIT_EMAIL=$${GIT_EMAIL} \
  	-e GIT_REPOSITORY=$${GIT_REPOSITORY} \
  	-e CODEGEN_SWAGGER_FILE=/swagger.yaml \
  	-e CODEGEN_LANGUAGE=go \
  	21stio/swagger-codegen-git-push