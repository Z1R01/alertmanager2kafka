FROM golang:1.21 as build

WORKDIR /go/src/github.com/webdevops/alertmanager2kafka

COPY ./go.mod /go/src/github.com/webdevops/alertmanager2kafka
COPY ./go.sum /go/src/github.com/webdevops/alertmanager2kafka
COPY ./Makefile /go/src/github.com/webdevops/alertmanager2kafka
COPY ./ /go/src/github.com/webdevops/alertmanager2kafka

RUN go mod download && go mod tidy
RUN make build
RUN ./alertmanager2kafka --help

FROM gcr.io/distroless/static
ENV LOG_JSON=1
COPY --from=build /go/src/github.com/webdevops/alertmanager2kafka/alertmanager2kafka /
USER 1000
ENTRYPOINT ["/alertmanager2kafka"]

