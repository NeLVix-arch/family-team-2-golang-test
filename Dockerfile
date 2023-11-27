FROM scratch

COPY server /

EXPOSE 80

CMD ["/server"]
