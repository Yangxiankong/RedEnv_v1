FROM ubuntu
WORKDIR /app
COPY . .
CMD [ "./REDENV_v1" ]