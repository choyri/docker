ARG VERSION=3.9.1


FROM rabbitmq:${VERSION}-alpine

ARG VERSION

LABEL maintainer="Chotow <i@choyri.com>"
LABEL version=${VERSION}

RUN cd /plugins; \
    wget https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.9.0/rabbitmq_delayed_message_exchange-3.9.0.ez; \
    rabbitmq-plugins enable --offline rabbitmq_delayed_message_exchange
