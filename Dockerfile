FROM golang:1.21.7 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# syntax=docker/dockerfile:1.3
FROM public.ecr.aws/lambda/go:1 AS base

# Set the platform for the build
FROM --platform=$BUILDPLATFORM baseWORKDIR /var/task

COPY --from=build-stage /main /main

CMD ["main"]
