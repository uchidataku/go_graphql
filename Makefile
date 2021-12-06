export

# ENV
DC = COMPOSE_FILE=docker-compose.yml docker-compose
DC_APP = $(DC) exec app

# コンテナ操作コマンド
build:
	$(DC) build
up:
	$(DC) up -d
build_up:
	$(DC) up -d --build
restart:
	$(DC) restart
force_restart:
	@make down
	@make build_up
down:
	$(DC) down
logs:
	$(DC) logs -f

run:
	$(DC_APP) go run main.go