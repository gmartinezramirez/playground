FROM python:3.9-slim-buster

COPY requeriments/requirements.txt /tmp/
COPY requeriments/requirements-testing.txt /tmp/
RUN pip install -r /tmp/requirements.txt
RUN pip install -r /tmp/requirements-testing.txt

RUN mkdir -p /src
COPY src/ /src/
RUN pip install -e /src
COPY tests/ /tests/

WORKDIR /src