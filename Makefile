.PHONY: run

run:
	@echo "Iniciando a aplicação..."
	docker-compose up -d

build:
	@echo "Building app"
	docker build -t zap-int:v0 .

docs:
	@echo "Updating swagger"
	swag init