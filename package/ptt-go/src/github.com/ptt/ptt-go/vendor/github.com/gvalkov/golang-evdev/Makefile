all: ecodes.go

HEADERS  = /usr/include/linux/input.h
HEADERS += /usr/include/linux/input-event-codes.h

ecodes.go: ecodes.go.template
	./bin/genecodes.py $(HEADERS) $< > $@

.PHONY: ecodes.go
