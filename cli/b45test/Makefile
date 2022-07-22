
all:
	go build

test1:
	go build
	mkdir -p test
	b45test --encode main.go --output ./test/out.out
	b45test --decode ./test/out.out --output ./test/out.pt2
	diff main.go ./test/out.pt2
	@echo PASS	
