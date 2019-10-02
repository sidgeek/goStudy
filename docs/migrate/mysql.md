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



CREATE TABLE IF NOT EXISTS users(
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  login_name VARCHAR (64) UNIQUE,
  pwd TEXT NOT NULL
);

UNSIGNED: 无符号
AUTO_INCREMENT： 自增
NOT NULL: 不为null
DEFAULT CURRENT_TIMESTAMP // 默认值
