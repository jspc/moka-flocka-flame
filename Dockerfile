FROM alpine
MAINTAINER jspc <james.condron@ft.com>

EXPOSE 80

ADD moka /moka
ADD json/ /json

CMD ["/moka", "-bind=:80"]
