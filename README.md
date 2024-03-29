# GoCommunicator

## Idea

Often I'm on a Linux, Windows or macOS device and I want some data from it.
Most of the time only some text but sometime some data too.

Here I want to build a light weight CLI application to upload these data to the specified server.

I like the idea from different system who use TXT-Records of domains.
The limit of 255 characters should be no problem.

```
example.org   IN   TXT   "gocom pk: [base64 public key]]"
example.org   IN   TXT   "gocom target: [https url for a http post]"
```

## Security

For Message Encryption is use the standard JWE with ECDH_ES_A256KW and the EC P-384 from NIST and A256GCM an Authenticated encryption (AE) algorithm. (Would like to use a different curve, than the NIST Curves.)
The target where the message is send (http post) to is read from a TXT Record so this could be spoofed.

## User Story

#### Use application on any maschine to send a message

![Console](https://imgur.com/OpqBIeb.png)

#### Get the message

Here I use the app Pushover (https://pushover.net/)

<img src="https://imgur.com/yJtgVXG.png" width="400">


## Sample Data

```
-----BEGIN PRIVATE KEY-----
MIGkAgEBBDDWDJb2FtQNmlPgR/IMvF8KpZAonNfCGQ0k/cZozfotWcw+ZyCY9Egu
SwsKWjkewqygBwYFK4EEACKhZANiAAQeybe7+OBzfa03Yutc+bf4cIvBQYFuE6ML
0pRJO2uW2V16AUpaj861yih/aCbE6DjmBvV10Nafar5Gdtvypsq+BZL9YjQIONCy
ezPjn6boRVm0Q+jIVLsoMWQt23077wg=
-----END PRIVATE KEY-----

 -----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEHsm3u/jgc32tN2LrXPm3+HCLwUGBbhOj
C9KUSTtrltldegFKWo/Otcoof2gmxOg45gb1ddDWn2q+Rnbb8qbKvgWS/WI0CDjQ
snsz45+m6EVZtEPoyFS7KDFkLdt9O+8I
-----END PUBLIC KEY-----
```
