# Compiler

FROM node:13.7.0-alpine AS builder
COPY package.json package-lock.json ./
RUN npm ci && mkdir /ng-app && mv ./node_modules ./ng-app
WORKDIR /ng-app
COPY . .
RUN npm run ng build -- --prod --output-path=dist

# Nginx

FROM nginx:1.17.8-alpine
#COPY nginx/default.conf.tpl /etc/nginx/conf.d/
COPY nginx/default.conf /etc/nginx/conf.d/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /ng-app/dist /usr/share/nginx/html
#CMD ["/bin/sh", "-c", "envsubst < /etc/nginx/conf.d/default.conf.tpl > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"]
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
