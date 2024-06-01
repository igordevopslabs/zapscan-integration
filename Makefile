.PHONY: run

run:
	@echo "Iniciando a aplicação..."
	CompileDaemon -command="./zapscan-integration" -color=true;

create-queue:
	@echo "Criando a fila sqs"
	aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name zap-queue
