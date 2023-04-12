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
- allow users to choose short url
- allow users to choose url ttl
- allow users to place number limits on short url accesses