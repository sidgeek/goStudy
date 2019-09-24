1、	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=example")

error: pq: SSL is not enabled on the server
解决：datasource = ”user=postgres password=*** host=localhost port=5432 dbname=testgo sslmode=disable“

参考：https://my.oschina.net/koalaone/blog/160165

2、gorm.Mode
// gorm.Model definition
type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}

CreatedAt: 创建时，自动写入
UpdatedAt: 更新时，自动写入
DeletedAt: 删除时，自动写入 （有了DeletedAt， table会执行软删除）

3、table
  1） Create Table (两种方法)
    db.AutoMigrate(&models.User{})
    db.CreateTable(&models.User{})
  2） Drop table
    db.DropTable(&models.User{})