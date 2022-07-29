
IN=/Volumes/x00/,in
QR=./out/tcs02.png

if [ -f ${QR} ] ; then 
	:
else
	./qr-secret --encode ${IN} --output ${QR}
	./qr-secret --decode ${QR} --output /tmp/,a
	cat /tmp/,a
	diff /tmp/,a ${IN}
	rm /tmp/,a
fi

