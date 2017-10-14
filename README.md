# Instructions

Install golang (Setup your GOPATH / GOROOT)

	wget http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz
	tar xzf ta-lib-0.4.0-src.tar.gz
	cd ta-lib
	./configure --prefix=/usr LDFLAGS="-lm"
	make
	make install
	go get -u github.com/d4l3k/talib
	go install github.com/d4l3k/talib

	go get -u github.com/matt-simons/ticker
