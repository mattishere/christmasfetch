PREFIX ?= /usr

install:
	@go build -o christmasfetch main.go
	@install -Dm755 christmasfetch $(DESTDIR)$(PREFIX)/bin/christmasfetch

uninstall:
	@rm -rf $(DESTDIR)$(PREFIX)/bin/christmasfetch
