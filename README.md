
# QR-Secret - A CLI tool for creating encrypted QR code data.

`qr-secret` is a command line tool for creating QR codes where the
data is AES256 bit encrypted.   You will be prompted for a password.

It can also take an image of a QR and decrypt that data so you can 
get back the original text file.

Data is limited to about 750 characters in a `256x256` bit QR code.

The command line functionality is in place --- but this is very much
a work in progress.

## Test

Run `make test` to run tests.

