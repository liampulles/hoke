FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY hoke .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/hoke"]