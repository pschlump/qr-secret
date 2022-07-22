
# QR-Secret - A CLI tool for creating encrypted QR code data.

```

-d "data" 	-
-i file		- read in a file
-o qr.png	- output a single qr code (Limits amount of data)
-m dir		- create multipe QR codes in dir
-p			- Password Prompt

-s inputQr.png - Input an Image to decode/decrypt
-t outputFile - file to put data into

3kb of data per QR code

http://qrcode.meetheed.com/question7.php

https://ts.q8s.co/ttt?d={...}
	{...} is the encrytped data

use base64url

github.com/dvsekhvalnov/jose2go/base64url

{"c":"...","d":"...","p":1,"q":14} base 64 URL encoded

1. DO Get Request
	- Get back Pub of Pub/Priv temporary key
	- Use Pub to encrypt data to send PW to it
	- Get back a "ID" to communicate with.
	https://ts.q8s.co/ttt?d={...}&pub=Caller's publick key - &id=XXX
	https://ts.q8s.co/ttt?d={...}&pw=Enc(pwd) &id=XXX
		-> returns text encrypte with key

```
