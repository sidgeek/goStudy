#使用CLI进行数据库(postgres)迁移

####参考：`base.md`
####区别：

###2）export database URL到一个变量中，方便后面操作使用
  ```
  export POSTGRESQL_URL=postgres://jack:password@localhost:5432/example?sslmode=disable
  ```

  ######其中：
  ```
  database: example
  user: jack
  password: password
  host: localhost
  ```

###3) 执行migrations指令
  ```
  migrate -database ${POSTGRESQL_URL} -path db/migrations up
  ```

