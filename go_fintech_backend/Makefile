# FILE Makefile BEGINS =======================================
include .env

# ANSI true color codes
# Netflix Red = 229,9,20
# Specific Blue = 3,219,248
# Bright Yellow = 255,255,0
RED=\033[38;2;229;9;20m
BLUE=\033[38;2;3;219;248m
YELLOW=\033[38;2;255;255;0m
NC=\033[0m # No Color
BOLD=\033[1m
RESET=\033[0m
# ====================================================|

define colorize_tests
	@echo "Running test in file $(1)_test.go..."
	@cd ./test && go test -v 2>&1 | \
	awk '/PASS/ {print "$(BOLD)$(BLUE)" $$0 "$(NC)"}; \
	/FAIL/ {print "$(BOLD)$(RED)" $$0 "$(NC)"}; \
	!/PASS|FAIL/ {print "$(YELLOW)" $$0 "$(NC)"}'
endef

# Function to print a bold message with lolcat
define print_bold_msg
	@echo "$(BOLD)$(1)$(RESET)" | lolcat
	@echo "$(RESET)"
endef

# Function to run a command in bold
define run_cmd
	@echo "$(BOLD)"; $(1) | lolcat; echo "$(RESET)"
endef
# ====================================================|

dev:
	$(call print_bold_msg,"Running go_fintech_backend project...")
	$(call run_cmd,go run ./cmd/main.go)
# ====================================================|

# creates a new migration
migrate:
	@(migrate create -ext sql -dir db/migrations/ -seq $(name)) 2>&1 | lolcat
# ====================================================|

# command to stop container
compose_down:
	@echo "Shutting down fintech_fullstack container..." | lolcat
	@docker compose down 2>&1 | lolcat

compose_up:
	@echo "running fintech_fullstack container..." | lolcat
	@docker compose up -d 2>&1 | lolcat

# running in detached mode: !!!CHECK DOCKER DESKTOP!!!
#create_container:
#	@docker run \
#          --name ${DOCKER_CONTAINER_NAME} \
#          -e POSTGRES_USER=${POSTGRES_USER} \
#          -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
#          -p 5432:5432 \
#          -d postgres:15.4-alpine3.18 | lolcat
# ====================================================|

db_up:
	@echo "Creating database..." | lolcat
	@docker exec -it fintech_fullstack createdb \
	--username=root --owner=root fintech_db

db_down:
	@echo "Dropping database..." | lolcat
	@docker exec -it fintech_fullstack dropdb \
	--username=root fintech_db
# ====================================================|

test1:
	$(call colorize_tests,mysql_connection)
# FILE Makefile END    =======================================

















