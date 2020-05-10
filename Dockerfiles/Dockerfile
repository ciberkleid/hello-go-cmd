FROM golang
LABEL com.example=hello
WORKDIR /workspace
COPY . .
RUN go install
ENTRYPOINT ["hello"]
