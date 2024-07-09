

.PHONY: 

all: sql build


sql: sqlc/*
	@ \
		cd sqlc; \
		echo "generating ORM..."; \
		sqlc generate;


build: cmd/* util/*
	@ \
		echo "building desktop_survivor..."; \
		CGO_ENABLED=0 GOOS=windows go build -o bin/desktop_survivor.exe cmd/main.go;


run: build
	@ ./bin/desktop_survivor


clean:
	@- rm -rf ui
	@- rm bin/desktop_survivor.exe
