FROM golang
LABEL maintainer="me@example.com"
WORKDIR /workspace
COPY . .
RUN go install
ENTRYPOINT ["hello"]
CMD ["world"]
