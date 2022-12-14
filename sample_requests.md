# Sample Requests

Some sample cURL requests are provided to help you interact with the API you build.  They default to `http://localhost:8000`.  You can also copy these commands into a REST client (e.g. Postman) if you prefer to use a GUI instead of the terminal.

## GET /ingest
```
curl --location --request POST 'http://localhost:8000/ingest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        "YYARKx|2022-03-22T21:42:02.362Z|-8",
        "YYARKx|2022-03-22T21:42:04.372Z|-1",
        "YYARKx|2022-03-22T21:42:50.572Z|7"
    ]
}'
```

## POST /device/{id}
```
curl --location --request GET 'http://localhost:8000/device/YYARKx'
```