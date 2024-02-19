#  TrueBackend Task

## Tech-stacks 
Development - `Golang` `Vue.Js` `SQLite DB` `Ant-vue`

Deployment - `Docker` `AWS` 
## Steps to configure backend locally

```
cd backend
go mod tidy
go run main.go
```
Run the following commands step by step

## Steps to configure backend locally

```
cd frontend
npm install
npm run serve
```
## Build using docker 

###  Backend
```
docker build -t techrounak/true-beacon-backend:v1 -f Backend.Dockerfile .
docker run -p 8000:8000 true-beacon-backend 
```
### Frontend
```
docker build -t techrounak/true-beacon-frontend:v1 -f Frontend.Dockerfile .
docker run -p 8080:8080 true-beacon-frontend 
```

### Ports 
Frontend - `localhost:8080`

Backend - `localhost:8000`