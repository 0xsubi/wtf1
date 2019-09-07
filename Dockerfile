FROM alpine

WORKDIR /app/

ADD views/* views/

COPY bin/wtf1 .

ENTRYPOINT [ "./wtf1" ]