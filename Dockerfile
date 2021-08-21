FROM debian:10-slim
COPY go-cli-example /
ENTRYPOINT [ "/go-cli-example" ]
