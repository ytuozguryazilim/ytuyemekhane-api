# ytuyemekhane-api ![Heroku](https://heroku-badge.herokuapp.com/?app=ytuyemekhane-api)
Web service of "ytu yemekhane" http://www.sks.yildiz.edu.tr/

# API - Routers

| Method  | Path                      | Description                                                                 |
| ------- |------------------------   |-----------------------------------------------------------------------------|
| GET     | /                         | Bugun olan oglen ve aksam yemegi listesi doner.                             |
| GET     | /:yil                     | Belli bir yildaki oglen ve aksam yemegi listesi doner.                      |
| GET     | /:yil/:ay                 | Belli bir yil ve aydaki oglen ve aksam yemegi listesi doner.                |
| GET     | /:yil/:ay/:gun            | Belli bir tarihteki oglen ve aksam yemegi listesi doner.                    |

# Localhost Kurulum

```bash
                   $ go get github.com/GnuYtuce/ytuyemekhane-api
                   $ cd $HOME/go/src/GnuYtuce/ytuyemekhane-api
(ytuyemekhane-api) $ # Manuel calistirma
(ytuyemekhane-api) $ PORT=8080; go run main.go
...
(ytuyemekhane-api) $ # Docker ile calistirma
(ytuyemekhane-api) $ sudo docker build -t ytuyemekhane-api .
(ytuyemekhane-api) $ sudo docker run -p 8080:8080 ytuyemekhane-api
...
```
