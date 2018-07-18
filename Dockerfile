FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build
CMD ["/usr/local/go/bin/go","/app/wp-attack"]
