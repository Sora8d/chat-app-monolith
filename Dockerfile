FROM golang:1.17

RUN mkdir monolith_api

COPY . monolith_api

WORKDIR monolith_api

RUN go get -d -v ./...

RUN go install -v ./...



CMD ["go", "run", "."]
