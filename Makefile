build:
	go build -o a.out \
		-ldflags "-X github.com/ydm/goinject/nope.Version=123 \
                          -X github.com/ydm/goinject.Version=234 \
                          -X github.com/ydm/goinject/package-with-dashes.Version=345 \
			  -X github.com/ydm/goinject/package-with-dashes/inner.Version=666 \
			  -X github.com/ydm/goinject/dashes-again.Version=456" \
		./cmd
