SHELL=sh -x

development:
	docker-compose run development-application

test:
	docker-compose run test-http-interface

publish_sdk:
	docker run \
		-v $$(pwd)/swagger.yaml:/swagger.yaml \
		-e NPM_PACKAGE_NAME=@21stio/ip-microservice-sdk \
		-e CODEGEN_SWAGGER_FILE=/swagger.yaml \
		-e CODEGEN_LANGUAGE=typescript-fetch \
		-e NPM_AUTH_TOKEN=$${NPM_AUTH_TOKEN} \
		-e NPM_EMAIL=$${NPM_EMAIL} \
		21stio/swagger-codegen-npm-publish