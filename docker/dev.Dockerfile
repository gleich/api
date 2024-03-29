FROM golang:1.15-alpine3.12

# Meta data
LABEL maintainer="email@mattglei.ch"
LABEL description="🛠  My personal GraphQL API"

# Copying over files
WORKDIR /usr/src/app
COPY . /usr/src/app

# Installing dependencies/
RUN go get -v -t -d ./...

# Setting up air
RUN go get -u github.com/cosmtrek/air
ENV air_wd /usr/src/app/

EXPOSE 80

CMD [ "air", "-c", ".air.toml" ]
