FROM nginx:stable-alpine
ADD recaptcha /bin/
RUN /bin/recaptcha > /var/log/recaptcha.log &
COPY server.conf /etc/nginx/conf.d/default.conf
COPY recaptcha.html /tmp/recaptcha.html
CMD  ["nginx", "-g" ,"daemon off;"]