Moka Flocka Flame
==

![waka-flocka-flame](http://news.hiphopearly.com/wp-content/uploads/2014/10/Waka-Flacka.png)


Building
--

```
CGO_ENABLED=0 go build -a -installsuffix cgo moka.go
```

Docker
--

Its a docker.

```
$ docker build -t moka .
$ docker run -P moka
```
