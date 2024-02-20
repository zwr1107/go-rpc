package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {
}

/*
1、方法只能有两个可序列化的参数，其中第二个参数是指针类型

	req 表示获取客户端传过来的数据
	res 表示给客户端返回数据

2、方法要返回一个error类型，同时必须是公开的方法。
3、req和res的类型不能是：channel（通道）、func（函数）均不能进行 序列化

//指针类型的参数是为了在方法内部改变参数的值
在 Go 语言中，指针是一个特殊的变量，它存储了另一个变量的内存地址。你可以通过指针来读取和修改这个内存地址上的数据，这就是指针的基本用途。  指针类型的声明方式是在类型前面加上一个星号（*）。例如，*int 表示一个指向整数的指针，*string 表示一个指向字符串的指针。
以下是一个简单的例子，演示了如何在 Go 语言中使用指针：
var x int = 1
var p *int = &x  // p 是一个指向整数的指针，&x 表示取 x 的内存地址
*p = 2           // *p 表示访问 p 指向的内存地址上的数据，这行代码将 2 赋值给 p 指向的内存地址上的数据
fmt.Println(x)   // 输出：2
在这个例子中，我们首先声明了一个整数变量 x 并赋值为 1，然后声明了一个指向整数的指针 p 并将 x 的内存地址赋值给 p。然后我们通过 *p 来修改 p 指向的内存地址上的数据，也就是修改 x 的值。最后，我们打印 x 的值，结果是 2，说明 x 的值确实被修改了
*/
func (this Hello) SayHello(req string, res *string) error {
	*res = "hello " + req
	return nil
}

func main() {
	//1.注册rpc服务 参数1：服务对象
	err1 := rpc.Register(new(Hello))
	if err1 != nil {
		panic(err1)
	}

	//2.监听服务
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	//3.应用退出的时候关闭监听
	defer listener.Close()

	//4.接收请求
	for {
		fmt.Println("服务启动" + listener.Addr().String())
		//建立连接
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//5.绑定服务
		go rpc.ServeConn(conn)
	}

}
