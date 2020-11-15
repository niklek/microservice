# modules image
FROM golang:1.15 AS modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# builder image
FROM golang:1.15 AS builder

COPY --from=modules /go/pkg /go/pkg

RUN useradd -u 10001 app
RUN mkdir -p /app
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 make build

# filan image
from scratch

COPY --from=builder /etc/passwd /etc/passwd
USER app

ARG PORT
ARG APP
ENV PORT ${PORT}
ENV APP ${APP}

COPY --from=builder /app/bin/${APP} /${APP}

EXPOSE ${PORT}

CMD ["/${APP}"]