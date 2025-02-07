# Dojo de GO
Este repositório de códigos desenvolvidos durante o treinamento da linguagem Go.

- GO:
go mod init "pasta-name"
go mod tidy

- DOCKER:
Criar arquivo dockerfile
- Dados:

   FROM mysql:8

   ENV MYSQL_ROOT_PASSWORD=root

   EXPOSE 3306
   CMD ["mysqld"]

   docker build -t mysql-dojo .
   docker run -d --name mysql-dojo -p 3306:3306 mysql-dojo

- MYSQL:
mysql -h 127.0.0.1 -P 3306 -u root -p
senha criada no docker MYSQL_ROOT_PASSWORD=root

- GORM:
go get -u gorm.io/gorm

- GO:
go get

Pesquisar:
air golang - stop app e roda de novo