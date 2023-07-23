FROM node:20-alpine3.17 as build
VOLUME /build/node_modules
WORKDIR /build
COPY package.json .
COPY yarn.lock .
RUN yarn
COPY . .
RUN yarn build

FROM nginx:1.25.1-alpine
COPY --from=build /build/dist/ usr/share/nginx/html/
