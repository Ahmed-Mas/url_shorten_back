# url_shorten_back
backend for url shortener app

This is going to be the backend for my URL shortener app. The app should recieve long URLs from the frontend, generate a short URL to be returned. Short URLs will be paired with long URLs using Redis as a temporary store.

general flow:
- storage <-> core-app <-> serializer <-> server <||> user

current capabilities:
- accept post calls with long urls
- generate short urls
- store "short:long" pairs in memory hash
- redirect short to long url

to-do:
- add redis as a "short:long" store
- add ttl to each generated short url
- allow users to choose short url (some char limit)
- allow users to choose url ttl (some max ttl)
- allow users to place number limits on short url accesses (1+ accesses, with 0 being unlimited)


Lets break down potential improvements to the code:
- core-app:
    - core-app struct:
        - connection to storage interface
    - methods:
        - generate short url
        - store short+long url
        - retrieve long url using short url
    - input/output struct:
        - long_url
        - short_url
        - timestamp
    - using 1 struct as input/output keeps things simple and makes it so it only 

- storage: - extendable to use different storage types
    - interface
        - Store
        - Retrieve

    - in-memory:
        - can only store data in a map
        - deletes all data after 10 entries

    - redis:
        - client

- serializer: - extendable to use different serializers
    - interface
        - Encode
        - Decode
    - turns incoming data into core-app struct
    - json:

- server:
    - router:
        - GET:
            - redirects to long url
        - POST:
            - generates short url
    