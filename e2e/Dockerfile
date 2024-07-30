FROM node:alpine

COPY . .

RUN npm i -y

ENTRYPOINT ["/bin/sh","-c","sleep 40 && npx playwright test"]