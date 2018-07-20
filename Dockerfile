FROM iron/go:dev
ADD . /go/src/github.com/hr1sh1kesh/vegeta-wp-load
ENV GOPATH /go

WORKDIR /go/src/github.com/hr1sh1kesh/vegeta-wp-load
RUN go build && mv vegeta-wp-load /usr/local/bin
CMD ["vegeta-wp-load"]