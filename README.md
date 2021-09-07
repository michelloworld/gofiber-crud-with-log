### Run without docker

run app

```
gin -a 3001 -p 3000
```

access by url

```
http://localhost:3000
```

### Run with docker (development)

```
docker-compose -f docker-compose.dev.yml up --build
```

### Curl for test


##### Create user

```
curl -X POST -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane.d@email.com"}' \
  http://localhost:3000/api/v1/users
```

###### Update user

```
curl -X PUT -H "Content-Type: application/json" \
  -d '{"name": "Hello World", "email": "hello.w@email.com"}' \
  http://localhost:3000/api/v1/users/1
```

##### Delete user

```
curl -X DELETE http://localhost:3000/api/v1/users/1
```