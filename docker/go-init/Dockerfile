FROM alpine:latest
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    git && \
    mkdir /src

WORKDIR /src
COPY ./cloner/cloner.sh /cloner.sh
RUN chmod a+x /cloner.sh

CMD ["/cloner.sh"]