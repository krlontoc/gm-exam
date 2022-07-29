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
Response:
```
{
    "data": [
        {
            "email_address": "kr@email.com",
            "full_name": "Kurt Russel"
        }
    ],
    "status": 200
}
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
{
    "email_address":"kr@email.com",
    "password":"Pass123!"
}
```
Response:
```
{
    "code": 200,
    "data": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6dHJ1ZSwic2Vzc2lvbiI6eyJpZCI6MSwiY3JlYXRlZF9hdCI6IjIwMjItMDctMjlUMTE6MTc6MTcuODU1NTE1OSswODowMCIsInVwZGF0ZWRfYXQiOiIyMDIyLTA3LTI5VDExOjE3OjE3Ljg1NTUxNTkrMDg6MDAiLCJmdWxsX25hbWUiOiJLdXJ0IFJ1c3NlbCIsImVtYWlsX2FkZHJlc3MiOiJrckBlbWFpbC5jb20ifSwiZXhwIjoxNjU5MDY2NDI0fQ.DIrKQBGpodI6CTfaohxqiyI0rG16UFYO1FivmfdqhvvoERqzzBwPZZipExzn264XguU6RGGDPsGohGwPg8rWpwRQoUsa5Ue5l-rs5WIH2Dxo-9CjARKOWK9E_jz2kMzaKFzx0AbLTlp1cVnSWufLsQeMdLS-W_bsHqalJINoRL4vFYy5H-9DZ6y28S26BkXFEvl4pDIjwpU6foPtS_OvjvenuvKMmFpO_CRgFkhVBLgbUYnFMOWjvoT38P_J4Cw6cFKmQzFjfPfo8ktfGCBixvEtOf0mFUV180H4HTPAMMQ85mX_3cMT93n2L0IK6xPSzoJHocSnPV5v_Ytb_wP_gA"
}
```
-- **/api/v1/user/{ID}** : get user info by id  
Method: `GET`  
URL: `/api/v1/user/<user id here>`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/user/1' \
--header 'Authorization: Bearer <token here>'
```
Response:
```
{
    "data": {
        "id": 1,
        "created_at": "2022-07-29T11:37:45.1469522+08:00",
        "updated_at": "2022-07-29T11:37:45.1469522+08:00",
        "full_name": "Kurt Russel",
        "email_address": "kr@email.com"
    },
    "status": 200
}
```
-- **/api/v1/user/balance** : get user balance  
Method: `GET`  
URL: `/api/v1/user/balance`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/user/balance' \
--header 'Authorization: Bearer <token here>'
```
Response:
```
{
    "data": 15000,
    "status": 200
}
```
-- **/api/v1/transactions** : get user transactions  
Method: `GET`  
URL: `/api/v1/transactions`  
cURL:
```
curl --location --request GET 'localhost:1007/api/v1/user/transactions' \
--header 'Authorization: Bearer <token here>'
```
Response:
```
{
    "data": [
        {
            "id": 4,
            "created_at": "2022-07-29T11:40:05.6085486+08:00",
            "updated_at": "2022-07-29T11:40:05.6085486+08:00",
            "from": 1,
            "to": 4,
            "amount": 2500,
            "message": "Allowance"
        },
        {
            "id": 3,
            "created_at": "2022-07-29T11:40:05.5937899+08:00",
            "updated_at": "2022-07-29T11:40:05.5937899+08:00",
            "from": 1,
            "to": 5,
            "amount": 2500,
            "message": "Allowance"
        },
        {
            "id": 2,
            "created_at": "2022-07-29T11:39:54.5022579+08:00",
            "updated_at": "2022-07-29T11:39:54.5022579+08:00",
            "from": 1,
            "to": 2,
            "amount": 5000,
            "message": "Rent Payment"
        },
        {
            "id": 1,
            "created_at": "2022-07-29T11:39:40.1977394+08:00",
            "updated_at": "2022-07-29T11:39:40.1977394+08:00",
            "from": 0,
            "to": 1,
            "amount": 10000,
            "message": "Deposit Transaction"
        }
    ],
    "status": 200
}
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
{
    "amount": 10000
}
```
Response:
```
{
    "data": {
        "id": 1,
        "created_at": "2022-07-29T11:39:40.1977394+08:00",
        "updated_at": "2022-07-29T11:39:40.1977394+08:00",
        "from": 0,
        "to": 1,
        "amount": 10000,
        "message": "Deposit Transaction"
    },
    "status": 200
}
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
{
    "to": "k@email.com",
    "amount": 5000,
    "message": "Rent Payment"
}
```
Response:
```
{
    "data": {
        "from": 1,
        "to": "k@email.com",
        "amount": 5000,
        "message": "Rent Payment"
    },
    "status": 200
}
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
Payload:
```
{
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
}
```
Response:
```
{
    "data": [
        {
            "send": {
                "to": "kl@email.com",
                "amount": 2500,
                "message": "Allowance"
            },
            "status": "OK",
            "title": "Success"
        },
        {
            "send": {
                "to": "rl@email.com",
                "amount": 2500,
                "message": "Allowance"
            },
            "status": "OK",
            "title": "Success"
        }
    ],
    "status": 200
}
```
