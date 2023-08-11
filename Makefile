.PHONY: all generate-data copy-db update-test-file copy-env

all: generate-data copy-db update-test-file copy-env

create-env:
	touch .env
	echo "API_PORT=22345" >> .env

generate-data:
	go run createInitDataInDB.go

copy-db:
	cp customer.db ./api/

copy-env:
	cp .env ./api/
