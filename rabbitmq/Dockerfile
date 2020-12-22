ARG BASE_IMAGE_TAG=alpine

FROM rabbitmq:${BASE_IMAGE_TAG}

LABEL maintainer="Chotow <i@choyri.com>"

RUN cd /plugins \
    && wget https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.8.9/rabbitmq_delayed_message_exchange-3.8.9-0199d11c.ez \
    && rabbitmq-plugins enable --offline rabbitmq_delayed_message_exchange
