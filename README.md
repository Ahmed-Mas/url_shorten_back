# url_shorten_back
backend for url shortener app

This is going to be the backend for my URL shortener app. The app should recieve long URLs from the frontend, generate a short URL to be returned. Short URLs will be paired with long URLs using Redis as a temporary store.

current capabilities:
- accept post calls with long urls
- generate short urls
- store "short:long" pairs in memory hash
- redirect short to long url

to-do:
- add redis as the "short:long" store
- add ttl to each generated short url
- allow users to choose short url (some char limit)
- allow users to choose url ttl (some max ttl)
- allow users to place number limits on short url accesses (1+ accesses, with 0 being unlimited)

issues:
- code separation of concerns isnt there yet
    - need to break out like so: storage <-> core-app <-> serializer <-> server <-> user
    - ^ this means having a redis pkg, core-app pkg, server pkg

Lets break down potential improvements to the code:
- core-app:
    - core-app struct:
        - connection to redis
    - methods:
        - generate short url
        - store short+long url
        - retrieve long url using short url
    - input/output struct:
        - long_url
        - short_url
        - timestamp
    - using 1 struct as input/output keeps things simple and makes it so it only 

- redis storage:
    - redis struct
        - redis client
    - methods:
        - store short+long url
        - retrueve long url using short url
    
    - input/output:
        - long_url
        - short_url
        - timestamp

- serializer:
    - maybe a serializer as a step between server and core-app would be useful
    - turns incoming json into core-app struct
    - methods:
        - Encode core-app struct into json bytes
        - Decode json bytes into core-app struct

- server:
    - router:
        - GET:
            - redirects to long url
        - POST:
            - generates short url
    