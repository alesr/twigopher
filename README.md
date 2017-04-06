# twigopher [![Build Status](https://travis-ci.org/alesr/twigopher.svg?branch=master)](https://travis-ci.org/alesr/twigopher)

This project integrates two different Go clients to access the Twitter streaming API and the accout API to stream tweets and get details related to the development API credentials.

After the initial implementation, a next feature would be to create a pool of twigopher clients using different credentials and runnining on Docker containers behind a load balancer and a RESTful API. Where the user would be able to send a new stream request and the application would use the available client/container accordingly with its credential usage.

The possibility of inserting or adding new credentials is also considered.

### HOWTO:

Add your Twitter credentials to the project root following the schema at `resources/test_credentials`

run `go run main.go -track "juventus" "golang"` to stream tweets containig those values.


@TODO:

 - More interfaces
 - Unit tests
