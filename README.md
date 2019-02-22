# Firebase as a microservice
An OMG service for Firebase, it allows to clod messaging with the subscribe client.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create Charge
```sh
$ omg run sendmessage -a topic=<TOPIC_NAME> -a body=<MESSAGE_BODY> -e SERVER_KEY=<SERVER_KEY>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-firebase .
```
### RUN
```
docker run -p 3000:3000 microservice-firebase
```