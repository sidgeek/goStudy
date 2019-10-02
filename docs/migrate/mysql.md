####参考：`base.md`
####区别：

###2）export database URL到一个变量中，方便后面操作使用
  ```
  export MYSQL_URL=mysql://user:password@tcp(host:port)/dbname?query
  ```

######例如
 ```
 export MYSQL_URL=mysql://root:123456@tcp(localhost:3306)/test
 ```

###3) 执行migrations指令
  ```
  migrate -database ${MYSQL_URL} -path migrations up
  ```