FROM golang:1.14-alpine AS build
MAINTAINER yeasy.github.com

WORKDIR /src/
COPY ./ /src/

RUN go mod tidy && \
	CGO_ENABLED=0 go build -o /bin/kafka-client

FROM scratch
COPY --from=build /bin/kafka-client /bin/kafka-client

ENTRYPOINT ["/bin/kafka-client"]