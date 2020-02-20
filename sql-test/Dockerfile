# Frontend

FROM node:13.7.0-alpine AS frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci && mkdir /ng-app && mv ./node_modules ./ng-app
WORKDIR /ng-app
COPY frontend/ .
RUN npm run ng build -- --prod --output-path=dist

# Backend

FROM golang:1.13.8 AS backend
WORKDIR /go/src/backend
COPY --from=frontend /ng-app/dist/. .
COPY backend/ .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sql-test .

# Runtime

FROM alpine:3.11.3 AS runtime
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=backend /go/src/backend/ .
EXPOSE 80
CMD ["./sql-test"]
