FROM ubuntu:latest
LABEL authors="sophis"

ENTRYPOINT ["top", "-b"]