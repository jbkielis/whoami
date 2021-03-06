FROM ubuntu:14.04

RUN apt-get update && apt-get install -y golang-go
RUN apt-get install bash
ADD . /app
WORKDIR /app
RUN go build -o http
ENV PORT 80
EXPOSE 80

CMD ["/app/http"]
