FROM node:18 AS build_stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

FROM node:18-alpine
WORKDIR /app

COPY --from=build_stage /app/package*.json ./
RUN npm install --production

COPY --from=build_stage /app/server.js .

EXPOSE 3001

CMD ["node", "server.js"]