FROM scratch

COPY service /service
COPY test.zip /test.zip

EXPOSE 80

CMD ["/service"]