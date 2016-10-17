# Hello Env

Really simple web server which reads an environment variable and says hello to that

Used by [swarm-setup](https://github.com/cookiefission/swarm-setup)

## Setup

    env GOOS=linux GOARCH=386 go build hello.go
    docker build -t seankenny/hello-env
