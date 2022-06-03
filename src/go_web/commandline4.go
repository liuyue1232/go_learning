package main
//头文件包导入
import (
    "fmt"//输出
    "flag"//flag包可以用来解析命令行参数
    "net"//用于tcp连接的包
    "sync"
    //"os"//os可以通过变量Args来获取命令参数，os.Args返回一个字符串数组
    "time"
)
//变量声明
var (
    c int
    n int
    s string
    wg   sync.WaitGroup//计数器，设置阻塞和计数
)

//从命令行读取参数
func init() {
    flag.IntVar(&c, "c", 0, "并发数")//变量名，输入参数名，默认值，帮助信息
    flag.IntVar(&n, "n", 0, "总请求数")
    flag.StringVar(&s, "s", "", "传送消息数据的内容")
    flag.Parse()//解析命令行参数，从os.Args[1:]开始。这个函数必须在所有的flag都定义好之后，程序获取flag之前调用。
}


//并发
func Concurrency() {
	var wg sync.WaitGroup
	var threeWg sync.WaitGroup//计数器，设置阻塞和计数,用于开多协程
	startTime := time.Now().Unix()
	for i := 0; i < c; i++ {//开c个协程并发
		wg.Add(c+1)
		go func(id int) {
                        fmt.Println(id, "线程发送") 
			for j := 0; j < n/c; j++ {//每个并发再开协程发送n/c条消息
				threeWg.Add(n/c+1)
				go func() {
                                        conn,err := net.Dial("tcp","1.116.92.132:8081")
                                        if err != nil{//直连失败报错
	                                    fmt.Println("客户端 dial err：",err)
	                                    return
                                        }
					_, err = conn.Write([]byte(s + "\n"))//发送消息
                                        
                                        if err != nil{
                                            fmt.Println("reconnect...",err)
                                        }//发送失败报错

				        conn.Close()
					

					threeWg.Done()
				}()
			}
			threeWg.Wait()
                        
			wg.Done()
		}(i)
	}
	wg.Wait()

	endTime := time.Now().Unix()
	elapsed := endTime - startTime
        fmt.Println("Finished in second. ",elapsed)
}

func main() {
    Concurrency()
    
    fmt.Println("-c",c)//打印输出
    fmt.Println("-n",n)
    fmt.Println("-s",s)
}
