FROM alpine:3.9

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD ./out/iap_auth .

ENV PORT=8081 \
	LOGGER_LEVEL=INFO \
  REFRESH_TIME_SECONDS= \
	IAP_HOST= \
	SERVICE_ACCOUNT_CREDENTIALS= \
	CLIENT_ID=
EXPOSE ${PORT}
CMD ./iap_auth