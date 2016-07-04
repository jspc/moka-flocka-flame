Moka Flocka Flame
==

![waka-flocka-flame](http://news.hiphopearly.com/wp-content/uploads/2014/10/Waka-Flacka.png)


Building
--

### Linux
```bash
$ CGO_ENABLED=0 go build -a -installsuffix cgo moka.go
```

### Non-linux
```bash
$ GOOS=linux GOARCH=amd64 go build moka
```

Docker
--

Its a docker.

```bash
$ docker build -t moka .
$ docker run -P moka
```

It also lives/berths at quay.io.

```bash
$ docker run -P quay.io/financialtimes/moka
```
