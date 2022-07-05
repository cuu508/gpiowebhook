# gpiowebhook


This is a tiny HTTP server for toggling Raspberry Pi GPIO pins via HTTP requests.

## Use

Rrun this on Raspberry Pi:

    KEY=secretkey ./gpiowebhook

This starts a HTTP server on port 28384. Now make a HTTP request like so:

    curl http://localhost:28384/?key=secretkey

The server will then set RPi's pin #17 to HIGH, wait 100ms, then set it to LOW.

Hook up something to the pins and you can now control it over HTTP.
In my case, I jerry-rigged a transmitter of an electric gate opener. The setup
will need some refinement :-)

![POC](/docs/poc.jpg?raw=true)

## Build

    env GOOS=linux GOARCH=arm go build
