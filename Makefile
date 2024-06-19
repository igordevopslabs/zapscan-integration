.PHONY: run

run:
	@echo "==> Iniciando a aplicação..."
	docker-compose up -d
	@echo "==> waiting app up..."

build:
	@echo "==> Building app"
	docker build -t zap-int:v0 .

docs:
	@echo "==> Updating swagger"
	swag init

install-tools: 
	@echo "==> Installing gotest"
	@go install github.com/rakyll/gotest@latest
	@echo "==> Installing swaggo"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "==> Installing staticcheck"
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@echo "==> Installing govulncheck"
	@go install golang.org/x/vuln/cmd/govulncheck@latest

go-checks:
	@echo "Rodando validações de segurança no codigo"
	@echo "==> Running staticcheck"
	@staticcheck ./...
	@echo "==> Running govulncheck"
	@govulncheck ./...

end-to-end:
	@echo "==> Running end-to-end tests"
	docker run --network="host" --rm -v $(PWD)/tests/e2e:/workdir --add-host=host.docker.internal:host-gateway jetbrains/intellij-http-client -L VERBOSE -e end_to_end -v http.client.env.json -r -D list.http 

run-e2e:
	@echo "==> Iniciando a aplicação para testes e2e..."
	docker-compose up -d
	@echo "==> waiting app up..."
	sleep 120
	docker-compose --profile=e2e up
	@echo "==> Finalizando stack da aplicação..."
	docker-compose down
	docker-compose --profile=e2e down