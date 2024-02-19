FROM node:18-alpine AS vue-build

WORKDIR /build/frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install
COPY frontend/ .
CMD ["npm", "run", "serve"]
EXPOSE 8080




