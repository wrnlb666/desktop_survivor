

.PHONY: 

all: sql

sql: sqlc/*
	@ \
		cd sqlc; \
		echo "generating ORM..."; \
		sqlc generate; \

clean:
	@- rm -rf ui
