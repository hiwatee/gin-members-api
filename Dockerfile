FROM golang:1.15.3

# locale & timezone (Asia/Tokyo)
# https://github.com/moby/moby/issues/12084
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# COPPY modulefile
COPY go.mod go.sum /go/src/
WORKDIR /go/src

# install goa
RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/jinzhu/gorm/dialects/postgres && \
    go get -u github.com/swaggo/swag/cmd/swag && \
    go get -u github.com/swaggo/gin-swagger && \
    go get -u github.com/swaggo/files
# install realize
RUN GO111MODULE=off go get github.com/oxequa/realize

# copy application code from host.
ADD . /go/src

# entrypoint permission settings
COPY ./docker/entrypoint.sh /
RUN chmod 777 /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]