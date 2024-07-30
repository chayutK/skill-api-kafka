FROM node:alpine

COPY . .

RUN npm i -y

ENTRYPOINT [ "npx","playwright","test" ]