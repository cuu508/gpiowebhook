# gpiowebhook


This is a tiny HTTP server for toggling GPIO pins via HTTP requests.
The current version calls `mraa-gpio` to do the toggling.

## Use

Run this on the SBC:

    KEY=secretkey ./gpiowebhook

This starts a HTTP server on port 28384. Now make a HTTP request like so:

    curl http://localhost:28384/?key=secretkey

The server will then set pin #11 to HIGH, wait 100ms, then set it to LOW.

Hook up something to the pins and you can now control it over HTTP.
In my case, I jerry-rigged a transmitter of an electric gate opener.

Proof-of-concept:

![POC](/docs/poc.jpg?raw=true)

Packaged up (the SBC I used is Radxa Zero):

![Case open](/docs/case-open.jpg?raw=true)

![Case closed](/docs/case-closed.jpg?raw=true)

## Build

    env GOOS=linux GOARCH=arm go build
