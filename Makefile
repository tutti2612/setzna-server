.PHONY: run
run:
	cp .realize.normal.yaml .realize.yaml
	docker-compose up -d

.PHONY: debug
debug:
	cp .realize.debug.yaml .realize.yaml
	docker-compose up -d

.PHONY: migration
migration:
	docker-compose exec app go run cmd/migration/migration.go

.PHONY: test
test:
	docker-compose exec app go test -v ./...