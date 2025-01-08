GOOSE_DRIVER ?= myqsl
GOOSE_DBSTRING ?= "root:sonha12@tcp(127.0.0.1:3306)/food_delivery_dev"
GOOSE_MIGRATION_DIR ?= ../../sql/schema

APP_NAME := server

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
run:
	go run main.go
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

upse_windows:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER) && set GOOSE_DBSTRING=$(GOOSE_DBSTRING) && set GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse_windows:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER) && set GOOSE_DBSTRING=$(GOOSE_DBSTRING) && set GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse_windows:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER) && set GOOSE_DBSTRING=$(GOOSE_DBSTRING) && set GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose -dir=$(GOOSE_MIGRATION_DIR) reset
statusse_windows:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER) && set GOOSE_DBSTRING=$(GOOSE_DBSTRING) && set GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: upse downse resetse run create_migration upse_windows downse_windows resetse_windowse statusse_windows
.PHONY: air