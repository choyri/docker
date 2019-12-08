FROM rabbitmq:alpine

LABEL maintainer="chotow <i@choyri.com>"

ENV DELAYED_MESSAGE_VERSION 3.8.0

RUN cd /plugins \
    && wget https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/v$DELAYED_MESSAGE_VERSION/rabbitmq_delayed_message_exchange-$DELAYED_MESSAGE_VERSION.ez \
    && rabbitmq-plugins enable --offline rabbitmq_delayed_message_exchange
