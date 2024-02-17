FROM alpine:3.5
MAINTAINER xljbear@gmail.com
WORKDIR /HoneyHollowServer

ENV MongoDBHost=127.0.0.1
ENV MongoDBPort=27017
ENV MongoDBUser=""
ENV MongoDBPassword=""
ENV MongoDBName=HoneyHollow
ENV ListenAddress=0.0.0.0:8080

COPY ./build/HoneyHollowServer_Linux/HoneyHollowServer /HoneyHollowServer
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata \
    && apk add --no-cache ca-certificates
RUN cd /HoneyHollowServer
ENTRYPOINT ./HoneyHollowServer
EXPOSE 8080