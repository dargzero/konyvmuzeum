FROM golang

WORKDIR /svc
COPY . .

RUN go install -v ./...

CMD ["konyvmuzeum"]