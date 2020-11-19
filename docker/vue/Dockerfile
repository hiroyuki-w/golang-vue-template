FROM node:lts-alpine3.12


WORKDIR /vue/src/app

RUN apk update && \
    npm install && \
    npm install -g npm && \
    npm install -g vue-cli

EXPOSE 8080
#####
#docker-compose exec vue sh
#vue init webpack
#npm install
#npm run dev