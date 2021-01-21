FROM golangci/golangci-lint:v1.33

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

# Set timezone to EST
ENV TZ=America/New_York
RUN echo "America/New_York" > /etc/timezone

EXPOSE 80

CMD [ "air", "-c", ".air.toml" ]
