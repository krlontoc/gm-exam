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

## Routes

## Testing
For local testing, use command `go test -v` to run the test.

## Links
The app is also upload to heroku, access the site using this [link](https://gm-exam.herokuapp.com/) `https://gm-exam.herokuapp.com`.
Get postman collection [here](https://www.getpostman.com/collections/2cdc08395f26e9047e4f) `https://www.getpostman.com/collections/2cdc08395f26e9047e4f`.
