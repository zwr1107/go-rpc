package main

import (
	"fmt"
	"net/rpc"
)

type Goods struct {
}

// AddGoodsReq AddGoods参数对应的结构体
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

// AddGoodsRes AddGoodsRes返回的结构体
type AddGoodsRes struct {
	Success bool
	Message string
}

// GetGoodsReq GetGoods参数对应的结构体
type GetGoodsReq struct {
	Id int
}

// GetGoodsRes GetGoodsRes返回的结构体
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	//rpc.Dial 连接rpc服务
	//conn 是一个客户端的连接
	//第一个参数是协议，第二个参数是服务端的地址
	conn, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	//关闭连接
	defer conn.Close()

	//定义返回值
	var res AddGoodsRes
	//调用远程方法，第一个参数是服务端的方法名，第二个参数是传递给服务端的参数，第三个参数是返回值
	err = conn.Call("Goods.AddGoods", AddGoodsReq{
		Id:      1,
		Title:   "我是客户端",
		Price:   12.5,
		Content: "我是客户端",
	}, &res)
	if err != nil {
		panic(err)
	}
	//打印结果
	fmt.Printf("%#v/n", res)

	//定义返回值
	var res2 GetGoodsRes
	//调用远程方法，第一个参数是服务端的方法名，第二个参数是传递给服务端的参数，第三个参数是返回值
	err = conn.Call("Goods.GetGoods", GetGoodsReq{
		Id: 1,
	}, &res2)

	if err != nil {
		panic(err)
	}

	//打印结果
	fmt.Printf("%#v", res2)
}
