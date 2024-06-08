## Fluxo para perform scan

### cadastrar app no tree via spider 
curl "http://localhost:8090/JSON/spider/action/scan/?apikey=fookey&url=http://juice-shop:3000"

### checar status - deve ser igual a 100 para seguir
curl "http://localhost:8090/JSON/spider/view/status/?apikey=fookey&scanId=0"

### Iniciar o scan ativo apos cadastro
curl "http://localhost:8090/JSON/ascan/action/scan/?apikey=fookey&url=http://juice-shop:3000"

### Verificar status do scan - deve ser igual a 100
curl "http://localhost:8090/JSON/ascan/view/status/?apikey=fookey&scanId=0"

### Report do scan 
curl "http://localhost:8090/JSON/ascan/view/scanProgress/?apikey=fookey&scanId=0"

## LocalHost

### cadastrar app no tree via spider 
curl "http://localhost:8090/JSON/spider/action/scan/?apikey=fookey&url=http://localhost:3000"

### checar status - deve ser igual a 100 para seguir
curl "http://localhost:8090/JSON/spider/view/status/?apikey=fookey&scanId=0"

### Iniciar o scan ativo apos cadastro
curl "http://localhost:8090/JSON/ascan/action/scan/?apikey=fookey&url=http://localhost:3000"

### Verificar status do scan - deve ser igual a 100
curl "http://localhost:8090/JSON/ascan/view/status/?apikey=fookey&scanId=0"

### Report do scan 
curl "http://localhost:8090/JSON/ascan/view/scanProgress/?apikey=fookey"