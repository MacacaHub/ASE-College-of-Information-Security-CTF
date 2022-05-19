FROM python:3.10-alpine3.14

RUN pip install ctfcli==0.0.10

COPY bootstrap.sh /

RUN chmod +x /bootstrap.sh

ENTRYPOINT ["/bootstrap.sh", "-d"]
