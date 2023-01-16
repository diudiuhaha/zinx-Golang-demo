package main

import "zinx/internal/zinx/znet"

/*
基于Zinx 框架来开发的服务器应用程序
*/
func main() {
	// 遇事不决，先写注释
	// 1. 创建一个server 句柄，即使用zinx的api
	s := znet.InitServer("[zinx V0.1]")
	// 2. 启动server
	s.Serve()
}
