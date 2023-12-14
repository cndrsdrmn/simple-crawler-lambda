.PHONY: build

build:
	sam build

invoke:
	sam local invoke --env-vars env.json

http:
	sam local start-api --env-vars env.json