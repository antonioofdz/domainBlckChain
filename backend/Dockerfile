FROM golang:alpine

WORKDIR /app/backend

RUN go get github.com/cespare/reflex

EXPOSE 8001

ENTRYPOINT reflex -r '\.go$' -s -- sh -c 'echo "Running local server" && go mod tidy && go run .'
