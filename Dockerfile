FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/simple-server /opt/bin/simple-server
ADD ./public /opt/public

CMD ["/opt/bin/simple-server"]
