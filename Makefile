.PHONY: push docker_build clean

push: docker_build
	docker push seankenny/hello-env

docker_build: hello
	docker build -t seankenny/hello-env .
	
hello: clean
	env GOOS=linux GOARCH=386 go build hello.go

clean:
	rm -f ./hello
