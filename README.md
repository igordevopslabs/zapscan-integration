# zapscan-intgegration
REST API for ZAProxy security scans


## Interagindo com a fila sqs
3. Publicar mensagens na fila:
```
aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/fake-queue --message-body "{\"namespace\": \"fake-inside-gw\",\"deployment\": \"fake-inside-gw\",\"action\": \"restart\"}"

#case2
aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/fake-queue --message-body "{\"namespace\": \"fake-inside-gw\",\"deployment\": \"fake-inside-gw\",\"action\": \"oomkill\"}"
```
4. Limpar a fila

```
aws --endpoint-url=http://localhost:4566 sqs purge-queue --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/minha-fila
```