func TestPrint(t *testing.T) {
  <!-- .... -->
}

1、写法： 必须以Test开头，并传入（t *testing.T）
t的参数类型：
  *testing.T 普通测试
  *testing.B Benchmark（性能）测试

2、跳过当前test
  func TestPrint(t *testing.T) {
    t.SkipNow() // 掉过当前test case,只能写在case开头
    ...
  }

3、子test
  go并不保证Test中的case按输书写顺序执行，如果需要顺序执行test，则应该使用subtests的写法
  func TestPrint(t *testing.T) {
    t.Run("a1", func(t *testing.T){fmt.Println("a1")})
    t.Run("a2", func(t *testing.T){fmt.Println("a2")})
    t.Run("a3", func(t *testing.T){fmt.Println("a3")})
  }

  <!-- 示例1 注意：以下写法，TestPrint1和TestPrint2都会执行两次 -->
  func TestPrint1(t *testing.T) {
    fmt.Println("testPrint")
  }

  func TestPrint2(t *testing.T) {
    fmt.Println("testPrint2")
  }

  func TestAll(t *testing.T) {
    t.Run("TestPrint1", TestPrint1)
    t.Run("TestPrint2", TestPrint2)
  }
  <!-- ******************************************************** -->

  <!-- 示例2 注意：(subtest开头T小写)以下写法，TestPrint1和TestPrint2只跑一次 -->
  func testPrint1(t *testing.T) {
    fmt.Println("testPrint")
  }

  func testPrint2(t *testing.T) {
    fmt.Println("testPrint2")
  }

  func TestAll(t *testing.T) {
    t.Run("TestPrint1", testPrint1)
    t.Run("TestPrint2", testPrint2)
  }
  <!-- ******************************************************** -->

4、所有Test的初始化(使用特殊函数TestMain)
  func TestMain(m *testing.M) {
    fmt.Println("test main first")
    m.Run() // 开始运行其它case,在它之前进行初始化
  }