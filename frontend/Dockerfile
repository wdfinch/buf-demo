FROM node:16-alpine
WORKDIR /app
COPY . .
RUN npm ci
RUN npm run build
FROM node:16-alpine
WORKDIR /app
COPY --from=0 /app .
CMD ["node", "./build"]