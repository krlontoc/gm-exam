## Description
This app is built in Golang ang SQLite (for data storage). Uses Iris for router handling and also for endpoint unit testing.

## How to run the app?
1. Build the docker image using this command.
```
docker build -t gm-exam
```
2. After a successful build, use this command to run the docker image.
```
docker run -d --rm -p 1007:1007 gm-exam
```

## Testing
For local testing, use command `go test -v` to run the test.

## Links
The app is also upload to heroku, access the site using this [link](https://gm-exam.herokuapp.com/) `https://gm-exam.herokuapp.com`.  
Get postman collection [here](https://www.getpostman.com/collections/2cdc08395f26e9047e4f) `https://www.getpostman.com/collections/2cdc08395f26e9047e4f`.

## Routes
- **/users** : get list of current users (for ease of use only)
Method: `GET`  
URL: `/users`  
cURL:
```
curl --location --request GET 'localhost:1007/users'
```
Payload:
```
```
Response:
```
```

- **/auth/v1/login** : get access token, token will be use for the other routes.
Method: `POST`  
URL: `/auth/v1/login`  
cURL:
```
curl --location --request POST 'localhost:1007/auth/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email_address":"kr@email.com",
    "password":"Pass123!"
}'
```
Payload:
```
```
Response:
```
```
-- **/api/v1/user/{ID}** : get user info by id  
Method: `GET`  
URL: `/api/v1/user/<user id here>`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/user/1' \
--header 'Authorization: Bearer <token here>'
```
Payload:
```
```
Response:
```
```
-- **/api/v1/user/balance** : get user balance  
Method: `GET`  
URL: `/api/v1/user/balance`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/user/balance' \
--header 'Authorization: Bearer <token here>'
```
Payload:
```
```
Response:
```
```
-- **/api/v1/transactions** : get user transactions  
Method: `GET`  
URL: `/api/v1/transactions`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/user/transactions' \
--header 'Authorization: Bearer <token here>'
```
Payload:
```
```
Response:
```
```
-- **/api/v1/transaction/depost** : depost amount 
Method: `POST`  
URL: `/api/v1/transaction/depost`  
cURL:
```
curl --location --request POST 'localhost:1007/api/v1/transaction/deposit' \
--header 'Authorization: Bearer <token here>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 10000
}'
```
Payload:
```
```
Response:
```
```
-- **/api/v1/transaction/send** : send amount to other user  
Method: `POST`  
URL: `/api/v1/transaction/send`  
cURL:
```
curl --location --request POST 'localhost:1007/api/v1/transaction/send' \
--header 'Authorization: Bearer <token here>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "to": "k@email.com",
    "amount": 5000,
    "message": "Rent Payment"
}'
```
Payload:
```
```
Response:
```
```
-- **/api/v1/transaction/multi-send** : send amount to multiple user at the same time   
Method: `POST`  
URL: `/api/v1/transaction/multi-send`  
cURL:
```
curl --location --request POST 'localhost:1007/api/v1/transaction/multi-send' \
--header 'Authorization: Bearer <token here>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sends": [
        {
            "to": "kl@email.com",
            "amount": 2500,
            "message": "Allowance"
        },
        {
            "to": "rl@email.com",
            "amount": 2500,
            "message": "Allowance"
        }
    ]    
}'
```
