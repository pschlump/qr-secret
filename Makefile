
all:
	go build


test: test001
	@echo PASS

test001:
	go build
	./qr-secret --encode ./testdata/t2 --output ./out/t2.png
	echo ""
	./qr-secret --decode ./out/t2.png --output ./out/t2.txt
	echo ""
	diff testdata/t2 out/t2.txt

install:
	rm -f ~/bin/qr-secret
	( cd ~/bin ; ln -s ../go/src/github.com/pschlump/qr-secret/qr-secret . )

