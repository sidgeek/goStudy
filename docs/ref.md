** https://awesome-go.com/#database

1、gorm: https://gorm.io/docs/index.html

2、migrate
 gormigrate: https://godoc.org/gopkg.in/gormigrate.v1
 golang-migrate/migrate: https://github.com/golang-migrate/migrate

postgres:
 https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

参考：https://blog.csdn.net/grace_yi/article/details/90243916

migrate CLI：https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
brew install golang-migrate

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=example sslmode=disable")

steps with cli:
1、(set the database URL to variable) When using Migrate CLI we need to pass to database URL. Let's export it to a variable for convienience:

For the purpose of this tutorial let's create PostgreSQL database called example. Our user here is postgres, password password, and host is localhost.

database: example
user: postgres
password: password
host: localhost

1、export POSTGRESQL_URL=postgres://postgres:password@localhost:5432/example?sslmode=disable
2、migrate create -ext sql -dir db/migrations -seq create_users_table
（在db/migrations生成两个文件）
000001_create_users_table.down.sql
000001_create_users_table.up.sql

3、In the .up.sql file let's create the table:
CREATE TABLE IF NOT EXISTS users(
   user_id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);

And in the .down.sql let's delete it:
DROP TABLE IF EXISTS users;

4、Run migrations
migrate -database ${POSTGRESQL_URL} -path db/migrations up

5、done, Let's check the table 





