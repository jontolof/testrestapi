# Dockerfile
# FROM iron/go
FROM scratch

WORKDIR /testrestapi

COPY main ./

EXPOSE 8080

ENTRYPOINT ["./main"]