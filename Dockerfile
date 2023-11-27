FROM scratch

COPY main /

EXPOSE 80

CMD ["/main"]
