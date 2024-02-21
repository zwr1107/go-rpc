package main

import "net/rpc"

func main() {
	//rpc.Dial 连接rpc服务
	//conn 是一个客户端的连接
	//第一个参数是协议，第二个参数是服务端的地址
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	//关闭连接
	defer conn.Close()

	//定义返回值
	var res string
	//调用远程方法，第一个参数是服务端的方法名，第二个参数是传递给服务端的参数，第三个参数是返回值
	err = conn.Call("Hello.SayHello", "我是客户端，你好", &res)
	if err != nil {
		panic(err)
	}
	//打印结果
	println(res)

}
