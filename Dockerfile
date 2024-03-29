FROM golang:latest 
ENV GOPATH=/
# WORKDIR /Users/evgenii/Desktop/GolangPAth/GoTestEducation/ypu
# RUN main --no-cache add bash git make gcc gettext
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

RUN go build -o ypu /Users/evgenii/Desktop/GolangPAth/GoTestEducation/ypu/main.go

CMD ["./ypu"]