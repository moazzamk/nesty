FROM scratch
COPY ./infrastructure/ca-certificates.crt /etc/ssl/certs/
ADD main /
CMD ["/main"]