GO=go

all: encode_decode test

test:
	$(GO) build test.go

encode_decode:
	$(GO) build encode_decode.go

run:
	./encode_decode test
	cmp test test.new && echo "test and test.new are same"

clean:
	-rm encode_decode
	-rm test
	-rm test.new
