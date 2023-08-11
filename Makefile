.PHONY: all create-env generate-data copy-db copy-env

all: create-env generate-data copy-db copy-env

create-env:
	touch .env
	echo "API_PORT=22345" >> .env

generate-data:
	go run createInitDataInDB.go

copy-db:
	cp customer.db ./api/

copy-env:
	cp .env ./api/
