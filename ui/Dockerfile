FROM node:22-slim AS base

WORKDIR /app

RUN apt-get update && apt-get upgrade -y --force-yes && mkdir -p ./app/public

COPY package.json package.json

FROM base AS prod

RUN npm i --only=production

COPY ./public /app/public

EXPOSE 3000

ENTRYPOINT ["npx", "http-server", "-p 3000"]

FROM base AS nonprod

RUN npm i

VOLUME ./app/public

EXPOSE 3000

ENTRYPOINT ["npx", "http-server", "-p", "3000"]

