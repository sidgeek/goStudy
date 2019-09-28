使用CLI进行数据库(postgres)迁移

推荐使用: golang-migrate/migrate
  https://github.com/golang-migrate/migrate


一、准备工作
1) 安装CLI（migrate CLI：https://github.com/golang-migrate/migrate/tree/master/cmd/migrate）
  brew install golang-migrate

2）export database URL到一个变量中，方便后面操作使用
  export POSTGRESQL_URL=postgres://jack:password@localhost:5432/example?sslmode=disable

  其中： database: example
        user: jack
        password: password
        host: localhost

二、表迁移
1）执行下面指令（在db/migrations生成两个文件）
  migrate create -ext sql -dir db/migrations -seq create_users_table

  生成
    000001_create_users_table.down.sql
    000001_create_users_table.up.sql

2）在两个文件中分别添加
  ...up.sql:
    CREATE TABLE IF NOT EXISTS users(
      user_id serial PRIMARY KEY,
      username VARCHAR (50) UNIQUE NOT NULL,
      password VARCHAR (50) NOT NULL,
      email VARCHAR (300) UNIQUE NOT NULL
    );

  ...down.sql:
    DROP TABLE IF EXISTS users;

3) 执行migrations指令
  migrate -database ${POSTGRESQL_URL} -path db/migrations up

4）检查数据表是否正确创建，结束。


三、其它
  参考: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md