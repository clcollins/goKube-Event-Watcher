FROM golang:1.12 AS build
ARG app_name
ENV appname=$app_name
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . /go/src/${appname}
WORKDIR /go/src/${appname}
RUN go get
RUN go build -o ${appname} -a -installsuffix cgo -v
RUN go test -v

FROM scratch
ARG app_name
ENV appname=$app_name
COPY --from=build /go/src/${appname}/${appname} /entrypoint
ENTRYPOINT ["/entrypoint"]
USER 1001
