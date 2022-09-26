FROM golang:1.15-alpine

COPY ./app /opt/app

WORKDIR  /opt

RUN echo "#!/bin/sh" > /bin/entrypoint \
    && echo "/opt/app" >> /bin/entrypoint \
    && chmod +x /bin/entrypoint

ENTRYPOINT ["entrypoint"]