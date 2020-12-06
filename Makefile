GO=go
GOBUILD=$(GO) build
BINDIR=build
BINCLI=kryptikcli
BIN=$(BINDIR)/$(BINCLI)
INSTALLLOC=/usr/local/bin/$(BINCLI)

all:
	$(GOBUILD) -o $(BIN)

clean:
	rm -rf $(BINDIR)/*

install:
	cp $(BIN) $(INSTALLLOC)

uninstall:
	rm $(INSTALLLOC)