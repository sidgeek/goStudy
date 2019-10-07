func TestPrint(t *testing.T) {
  <!-- .... -->
}

1. 写法： 必须以Test开头，并传入（t *testing.T）

   ######t的参数类型：
     * *testing.T 普通测试
     * *testing.B Benchmark（性能）测试

2. 跳过当前test
    ```
    func TestPrint(t *testing.T) {
        t.SkipNow() // 掉过当前test case,只能写在case开头
        ...
      }
    ```

3. 子test

      go并不保证Test中的case按输书写顺序执行，如果需要顺序执行test，则应该使用subtests的写法
      ```
      func TestPrint(t *testing.T) {
        t.Run("a1", func(t *testing.T){fmt.Println("a1")})
        t.Run("a2", func(t *testing.T){fmt.Println("a2")})
        t.Run("a3", func(t *testing.T){fmt.Println("a3")})
      }
      ```

      ####示例1 注意：以下写法，TestPrint1和TestPrint2都会执行两次
      ```
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
      ```
    
      ####示例2 注意：(subtest开头T小写)以下写法，TestPrint1和TestPrint2只跑一次
      ```
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
      ```

4. 所有Test的初始化(使用特殊函数TestMain)

      ######参数类型: *testing.M
      ```
      func TestMain(m *testing.M) {
        fmt.Println("test main first")
        m.Run() // 开始运行其它case,在它之前进行初始化
      }
      ```

5. benchmark

    * case 以Benchmark开头
    * 参数类型: *testing.B
    * benchmark的case一般会跑b.N次，而且每次执行都会如此
    * 在执行过程中会根据实际case的执行时间是否稳定会增加b.N次数以达到稳态

    ######执行:
      `go test -bench=.`
    
    ######参考：
      `https://cloud.tencent.com/developer/article/1469185`

6. 测试web框架的性能（go）:
    * https://github.com/smallnest/go-web-framework-benchmark
    * docker 版本： https://hub.docker.com/r/smallnest/go-web-framework-benchmark/

7. 对不同语言，不同框架的web测试
    * https://github.com/mihaicracan/web-rest-api-benchmark

8. 压力测试工具-- wrk
    * rew install wrk

9. 注意
    * 测试单个文件时，注意需要指定测试文件的所有依赖，如
    ````
    go test -v api/dbops/api_test.go api/dbops/api.go api/dbops/conn.go
    ````
   其中，api.go的运行依赖了conn.go，则需要加上


