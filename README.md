# _Firebase_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/firebase.svg?branch=master)](https://travis-ci.com/omg-services/firebase)
[![codecov](https://codecov.io/gh/omg-services/firebase/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/firebase)


An OMG service for Firebase, it allows to clod messaging with the subscribe client.

## Direct usage in [Storyscript](https://storyscript.io/):
##### Send Message By Token
```coffee
>>> firebase sendMessageByToken token:'validToken' title:'messageTitle' body:'messageBody' data:'{"data":"object"}'
{"Ok": true,"StatusCode": 200,"multicast_id": 5795905028209682000,"success": 1,"failure": 0,"canonical_ids": 0,"results": [{"message_id": "0:1560855298038939%2fd9afcdf9fd7ecd"}], "RetryAfter": ""}
```
##### Send Message By Topic
```coffee
>>> firebase sendMessageByTopic token:'validToken' topic:'messageTopic' body:'messageBody' data:'{"data":"object"}'
{"Ok": true,"StatusCode": 200,"multicast_id": 7095436517463231000,"success": 1,"failure": 0,"canonical_ids": 0,"results": [{"message_id": "0:1560855941969998%e609af1cf9fd7ecd"}], "RetryAfter": ""}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Send Message By Token
```sh
$ omg run sendMessageByToken -a token=<TOKEN> -a title=<NOTIFICATION_TITLE> -a body=<NOTIFICATION_BODY> -a icon=<NOTIFICATION_ICON> -a data=<DATA_OBJECT>  -e SERVER_KEY=<SERVER_KEY>
```
##### Send Message By Topic
```sh
$ omg run sendMessageByTopic -a token=<TOKEN> -a topic=<TOPIC> -a body=<NOTIFICATION_BODY> -a icon=<NOTIFICATION_ICON> -a data=<DATA_OBJECT>  -e SERVER_KEY=<SERVER_KEY>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/firebase/blob/master/LICENSE).
