## What's this?

This is a modified version of [go blog api](https://github.com/abdurraufraihan/golang-blog-api).
For demo only.

## Run locally

```bash
docker-compose up --build
```

## Test
```bash

http GET http://localhost:8000/api/v1/contacts 

http POST http://localhost:8000/api/v1/contacts name=Arief  email=ariefh@unknown.com \
mobile_no=+6281363531111 institution=Unknown

http GET http://localhost:8000/api/v1/contacts 

ID=$(curl http://localhost:8000/api/v1/contacts | jq -r '.data[-1].id')
http PUT http://localhost:8000/api/v1/contacts/$ID name="Arief Hidayat" email=ariefh1@unknown.com mobile_no=+6586683317 institution=Unknown

```

## Clean up
```bash
docker-compose down
```