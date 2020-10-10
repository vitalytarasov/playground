APP = play
IMAGE = play

# builds executable for container
build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o $(APP)

# removes executable and containers
clean:
	rm -f $(APP)
	docker rm -f $(IMAGE)

# builds docker image
image: build
	docker build -t $(IMAGE) .

# runs container
run: image
	docker run -t $(IMAGE) $(IMAGE)
