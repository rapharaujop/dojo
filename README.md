# Dojo de GO
Este repositório de códigos desenvolvidos durante o treinamento da linguagem Go.

1. **GO:**
go mod init "pasta-name"
go mod tidy

2. **DOCKER:**
Criar arquivo dockerfile
- Dados:

   FROM mysql:8

   ENV MYSQL_ROOT_PASSWORD=root

   EXPOSE 3306
   CMD ["mysqld"]

   docker build -t mysql-dojo .
   docker run -d --name mysql-dojo -p 3306:3306 mysql-dojo

3. **MYSQL:**
mysql -h 127.0.0.1 -P 3306 -u root -p
senha criada no docker MYSQL_ROOT_PASSWORD=root

4. **GORM:**
go get -u gorm.io/gorm

5. **GO:**
go get