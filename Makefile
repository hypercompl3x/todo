.PHONY: tailwind-watch
tailwind-watch:
		./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

.PHONY: tailwind-build
tailwind-build:
		./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

.PHONY: templ-generate
templ-generate:
		templ generate

.PHONY: dev
dev:
		go build -o ./tmp/main.exe ./cmd/main.go && air

.PHONY: build
build:
		make tailwind-build
		make templ-generate
		go build -o ./tmp/main.exe ./cmd/main.go