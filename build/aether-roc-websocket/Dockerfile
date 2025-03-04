# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0.0 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/aether-roc-api/
COPY cmd/ /go/src/github.com/onosproject/aether-roc-api/cmd/
COPY pkg/ /go/src/github.com/onosproject/aether-roc-api/pkg/
COPY build/ /go/src/github.com/onosproject/aether-roc-api/build/
COPY vendor/ /go/src/github.com/onosproject/aether-roc-api/vendor/

RUN cd /go/src/github.com/onosproject/aether-roc-api && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

FROM alpine:3.12
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/onosproject/aether-roc-api/build/_output/aether-roc-websocket /usr/local/bin/aether-roc-websocket

ENTRYPOINT ["aether-roc-websocket"]
