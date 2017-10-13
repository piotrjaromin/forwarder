FROM scratch

WORKDIR  /app

ADD bin/forwarder .

ENTRYPOINT ["./forwarder"]