
PATH_TO_TEST=tests/e2e

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
	@DOCKER_INTERNAL_IP=$$(docker network inspect bridge | grep Gateway | awk '{print $$2}' | tr -d '",') && \
	docker run --network="host" --rm -v $(PWD)/$(PATH_TO_TEST):/workdir --add-host=host.docker.internal:$$DOCKER_INTERNAL_IP jetbrains/intellij-http-client -L VERBOSE -e end_to_end -v http.client.env.json -r -D list.http

