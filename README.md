# ytuyemekhane-api ![Heroku](https://heroku-badge.herokuapp.com/?app=ytuyemekhane-api)
Web service of "ytu yemekhane" http://www.sks.yildiz.edu.tr/

# API - Routers

| Method  | Path                      | Description                                                                 |
| ------- |------------------------   |-----------------------------------------------------------------------------|
| GET     | /                         | Bugun olan oglen ve aksam yemegi listesi doner.                             |
| GET     | /:yil                     | Belli bir yildaki oglen ve aksam yemegi listesi doner.                      |
| GET     | /:yil/:ay                 | Belli bir yil ve aydaki oglen ve aksam yemegi listesi doner.                |
| GET     | /:yil/:ay/:gun            | Belli bir tarihteki oglen ve aksam yemegi listesi doner.                    |
