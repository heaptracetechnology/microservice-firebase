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

##### Send Message By Token
```sh
$ omg run send_message_by_token -a token=<TOKEN> -a title=<NOTIFICATION_TITLE> -a body=<NOTIFICATION_BODY> -a icon=<NOTIFICATION_ICON> -a data=<DATA_OBJECT>  -e SERVER_KEY=<SERVER_KEY>
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