This is a tiny HTTP server for toggling Raspberry Pi GPIO pins via HTTP requests.

Usage (run this on Raspberry Pi):

    KEY=secretkey ./gpiowebhook

This starts a HTTP server on port 28384. Now make a HTTP request like so:

    curl http://localhost:28384/?key=secretkey

The server will then set RPi's pin #17 to HIGH, wait 100ms, then set it to LOW.

And that's all :-)

Build it for 32-bit ARM:

    env GOOS=linux GOARCH=arm go build

