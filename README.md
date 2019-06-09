# GoCommunicator

## Idea

Often I'm on a Linux, Windows or macOS device and I want some data from it.
Most of the time only some text but sometime some data too.

Here I want to build a light weight CLI application to upload these data to the specified server.

I like the idea from different system who use TXT-Records of domains.
The limit of 255 characters should be no problem.

```
goco.example.org   IN   TXT   "pk: [base64 public key], url: [https url for a http post]"
```