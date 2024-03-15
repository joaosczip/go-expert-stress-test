# Go Expert Stress Test

Este projeto é uma ferramenta de linha de comando escrita em Go que realiza testes de stress em uma URL específica.

### Construindo a imagem Docker

Para construir a imagem Docker deste projeto, você precisa ter o Docker instalado em sua máquina. Depois de instalado, você pode construir a imagem usando o seguinte comando:

```sh
docker build -t nome-da-sua-imagem .
```

Substitua nome-da-sua-imagem pelo nome que você deseja dar à sua imagem Docker.

### Executando o comando de stress test

Depois de construir a imagem Docker, você pode executar o comando de stress usando o seguinte comando:

```sh
docker run nome-da-sua-imagem stress --url=http://google.com --requests=1000 --concurrency=10
```

Substitua `nome-da-sua-imagem` pelo nome que você deu à sua imagem Docker. Você pode substituir `http://google.com` pela URL que deseja testar. `--requests=1000` especifica o número de solicitações a serem feitas e `--concurrency=10` especifica o número de solicitações simultâneas.

### Exemplos

```
# Realizar 1000 solicitações para http://google.com com 10 solicitações simultâneas
docker run nome-da-sua-imagem stress --url=http://google.com --requests=1000 --concurrency=10

# Realizar 5000 solicitações para http://example.com com 50 solicitações simultâneas
docker run nome-da-sua-imagem stress --url=http://example.com --requests=5000 --concurrency=50

# Realizar 100 solicitações para http://localhost:8080 com 5 solicitações simultâneas
docker run nome-da-sua-imagem stress --url=http://localhost:8080 --requests=100 --concurrency=5
```

Lembre-se de substituir `nome-da-sua-imagem` pelo nome que você deu à sua imagem Docker.