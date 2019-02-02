# Build command:
# docker build -t bernardosecades/warning-shot .
# Run command:
# docker run -td --name warning_shot -p 3000:3000 bernardosecades/warning-shot
FROM        golang:1.11
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Install needed packages for container
RUN         apt-get update --fix-missing && apt-get upgrade -y && apt-get install -y \
            git \
            vim

# Force the go compiler to use modules
ENV         GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files
RUN         mkdir /app
WORKDIR     /app
COPY        go.mod .
COPY        go.sum .

# Download all the dependencies that are specified in the go.mod and go.sum file
RUN         go mod download

# build directories and build
ADD         . /app/
RUN         go build

# Expose ports and run
EXPOSE      3000
CMD         ["./warning-shot"]
