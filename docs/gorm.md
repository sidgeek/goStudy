参考: 
  https://gorm.io/docs/index.html
  https://jasperxu.github.io/gorm-zh/ 中文文档

1、gorm.Mode
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