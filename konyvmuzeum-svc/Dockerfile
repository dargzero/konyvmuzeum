FROM golang

WORKDIR /konyvmuzeum-svc
COPY . .

RUN go install -v ./...

CMD ["konyvmuzeum"]