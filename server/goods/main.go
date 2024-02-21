package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
}

// AddGoods参数对应的结构体
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}
type AddGoodsRes struct {
	Success bool
	Message string
}

// GetGoods参数对应的结构体
type GetGoodsReq struct {
	Id int
}
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (this Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	//1、执行增加 模拟
	fmt.Printf("%#v", req)
	//2、返回增加的结果
	*res = AddGoodsRes{
		Success: true,
		Message: "增加数据成功",
	}
	return nil
}

func (this Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	//1、执行增加 模拟
	fmt.Printf("%#v", req)
	//2、返回增加的结果
	*res = GetGoodsRes{
		Id:      12,
		Title:   "服务器获取的数据",
		Price:   24.5,
		Content: "我是服务器数据库获取的内容",
	}
	return nil
}

type GoodsInfo struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	//1.注册rpc服务 参数1：服务对象
	err1 := rpc.Register(new(Goods))
	if err1 != nil {
		panic(err1)
	}

	//2.监听服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("服务启动" + listener.Addr().String())

	//3.应用退出的时候关闭监听
	defer listener.Close()

	//4.接收请求
	for {
		//建立连接
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//5.绑定服务
		go rpc.ServeConn(conn)

		fmt.Println("服务启动" + listener.Addr().String())
	}
}
