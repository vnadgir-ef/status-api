FROM scratch

ADD status-api /status-api
ADD config.yml /config.yml

CMD ["./status-api", "-configFile=config.yml"]
