.PHONY:
REPO=cmattoon
IMAGE_NAME=aws-env
IMAGE=$(REPO)/$(IMAGE_NAME)

container:
	docker build -t $(IMAGE):latest .

push: container
	docker push $(IMAGE):latest
