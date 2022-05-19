FROM node:16-alpine3.14

WORKDIR /app
COPY ./app .

RUN npm install

EXPOSE 3000

CMD ["node", "server.js"]
