FROM scratch
WORKDIR /app
ADD . /app/
EXPOSE 8001

CMD [ "./main"]