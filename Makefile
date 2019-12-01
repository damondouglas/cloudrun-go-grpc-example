export project=$$(gcloud config get-value project)
export image=gcr.io/${project}/echo:1.0

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen: ## Stub proto
	protoc -I proto proto/echo.proto --go_out=plugins=grpc:proto

build: ## Build docker image
	docker build -t ${image} .

sh: ## sh into image
	docker run -it ${image} sh

run: ## run image locally
	docker run -it -p 8080:8080 ${image}

push: ## push image
	gcloud auth configure-docker; \
	docker push ${image}

deploy: ## Deploy cloud run service
	gcloud run deploy echo --image="${image}" --allow-unauthenticated --project ${project} --region us-central1 --platform managed

print-host: ## Print cloud run host endpoint
	@gcloud run services list --platform managed | grep https | tr -s " " | cut -f 4 -d " "