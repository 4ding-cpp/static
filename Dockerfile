FROM nginx:stable-alpine
COPY server.conf /etc/nginx/conf.d/default.conf
CMD  ["nginx", "-g" ,"daemon off;"]