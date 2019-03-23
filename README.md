# ytuyemekhane-api
[![MicroBadger Size](https://img.shields.io/microbadger/image-size/image-size/ef9n/ytuyemekhane-api.svg)](https://hub.docker.com/r/ef9n/ytuyemekhane-api)
[![License (GPL version 3)](https://img.shields.io/badge/license-GNU%20GPL%20version%203-red.svg?style=flat-square)](https://www.gnu.org/licenses/gpl-3.0.en.html)
[![Go Report Card](https://goreportcard.com/badge/GnuYtuce/ytuyemekhane-api)](https://goreportcard.com/report/GnuYtuce/ytuyemekhane-api)
[![GoDoc](https://godoc.org/github.com/GnuYtuce/ytuyemekhane-api?status.svg)](https://godoc.org/github.com/GnuYtuce/ytuyemekhane-api/)
[![Heroku](https://heroku-badge.herokuapp.com/?app=ytuyemekhane-api)](https://ytuyemekhane-api.herokuapp.com/)

Web service of "ytu yemekhane" http://www.sks.yildiz.edu.tr/

# API - Routers

| Method | Path           | Description                                                  |
| ------ | -------------- | ------------------------------------------------------------ |
| GET    | /              | Bugun olan oglen ve aksam yemegi listesi doner.              |
| GET    | /:yil          | Belli bir yildaki oglen ve aksam yemegi listesi doner.       |
| GET    | /:yil/:ay      | Belli bir yil ve aydaki oglen ve aksam yemegi listesi doner. |
| GET    | /:yil/:ay/:gun | Belli bir tarihteki oglen ve aksam yemegi listesi doner.     |

# Docker ile Kurulum

```bash
$ docker image pull ef9n/ytuyemekhane-api
$ docker container run -p 8080:8080 ef9n/ytuyemekhane-api 
```
