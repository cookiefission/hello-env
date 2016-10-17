FROM alpine
MAINTAINER Sean Kenny

EXPOSE 8000
CMD ["hello"]

COPY hello /usr/local/bin/hello
RUN chmod +x /usr/local/bin/hello
