FROM admxj/go-beego:v0.0.1 as builder

COPY . /code
WORKDIR /code
ARG envType

RUN bee pack -be GOOS=linux GOARCH=amd64

FROM admxj/centos:v0.0.2
RUN mkdir /www
COPY --from=builder /code/code.tar.gz /www
RUN tar -xvf /www/code.tar.gz -C /www
RUN rm -rf /www/code.tar.gz
CMD ["/www/code"]