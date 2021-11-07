FROM ubuntu
WORKDIR /app
COPY . .
ENV MYSQL_HOST="value"
ENV MYSQL_USER="root"
ENV MYSQL_PASSWORD="Syhyeyh95588"
ENV REDIS_HOST="redis-cn02y38m7f5ehl1xs.redis.ivolces.com:6379"
ENV REDIS_PASSWORD="Syhyeyh95588"
CMD [ "./REDENV_v1" ]