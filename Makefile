APP = play
IMAGE = play

# builds executable for alpine container
build:
	GOOS=linux CGO_ENABLED=0 go build -o $(APP)

# removes executable, container and image
clean:
	-rm -f $(APP)
	-docker container rm $(IMAGE)
	-docker rmi -f $(IMAGE)

# formats go sources
fmt:
	go fmt playground/...

# builds docker image
image: build
	docker build -t $(IMAGE) .

# runs container
run: image
	docker run --name $(IMAGE) -p 1080:1080 $(IMAGE)
