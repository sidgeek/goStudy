1、	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=example")

error: pq: SSL is not enabled on the server
解决：datasource = ”user=postgres password=*** host=localhost port=5432 dbname=testgo sslmode=disable“

参考：https://my.oschina.net/koalaone/blog/160165


