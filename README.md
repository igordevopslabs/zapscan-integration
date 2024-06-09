# ZapScan Integration

---

## Breve Descrição do Projeto

O **ZapScan Integration** é um projeto para fins de estudo desenvolvido para automatizar o processo de varreduras de segurança em aplicações web utilizando o OWASP ZAP (Zed Attack Proxy). Este projeto fornece uma API REST que se comunica com uma instância do ZAP, permitindo a execução de varreduras de segurança e a consulta dos resultados dessas varreduras. O objetivo principal é demonstrar como integrar o ZAP em um ambiente controlado e gerenciar varreduras de segurança de forma automatizada.

---

## Ferramentas Mínimas e Requisitos

Para executar este projeto, você precisará das seguintes ferramentas e versões mínimas:

- **Golang**: 1.21 ou superior
- **Docker**: 20.10 ou superior
- **Docker Compose**: 1.29 ou superior
- **Make**: Ferramenta para automação de build

---

## Como Rodar o Projeto?

Aqui estão as instruções passo a passo para configurar e rodar o projeto:

1. **Clone o Repositório**

   Clone o repositório do projeto para a sua máquina local:

   ```sh
   git clone https://github.com/igordevopslabs/zapscan-integration.git
   cd zapscan-integration
   ```
2. **Instale as Ferramentas Necessárias**

   Certifique-se de ter as ferramentas mencionadas nos requisitos instaladas. 
   Você pode usar o comando make install-tools para instalar algumas ferramentas Go adicionais necessárias para 
   este projeto.

3. **Compile o Projeto**

   Compile o projeto usando o Docker:
   ```sh
   make build
   ```

4. **Inicie a stack**

   ```sh
   make run
   ```
   Isso irá iniciar todos os serviços definidos no arquivo docker-compose.yml.

## Action e Testes de Segurança
O projeto inclui uma GitHub Action configurada para executar verificações de segurança no código Go sempre que houver mudanças no repositório. A configuração da action está no arquivo workflow.yml e executa as seguintes etapas:

  * Linting: Verifica o código Go utilizando ferramentas como staticcheck e govulncheck.

Para rodar os testes de segurança manualmente, utilize o comando:
```sh
make go-checks
```
## Acessando a documentação da API.

Esse projeto possui um endpoint com a documentação da API (Swagger.)

Você pode encontrá-lo através da rota ```/swagger/index.html```

Se você precisar atualizar a documentação Swagger, utilize o comando:
```sh
make docs
```

## Aviso de Uso Acadêmico

**Aviso**: Este projeto é destinado exclusivamente para fins acadêmicos e de aprendizado. Ele não deve ser utilizado em ambientes de produção, pois não possui garantias de segurança e estabilidade para tais fins.