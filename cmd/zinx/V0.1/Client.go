package main

import (
	"log"
	"net"
	"time"
)

/*
	模拟客户端
*/

func main() {
	log.Println("client start...")
	// 1. 直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("connect server failed ,err :", err)
		return
	}
	// 2. 连接调用write 写数据,做循环
	for {
		_, err := conn.Write([]byte("hello world!")) // 写入信息
		if err != nil {
			log.Println("write failed ,err:", err)
			return
		}
		msg := make([]byte, 512)
		read, err := conn.Read(msg) // 读取信息
		if err != nil {
			log.Println("read msg failed ,err :", err)
			return
		}

		log.Printf("server call back:%s ,read = %d \n", msg, read)
	}
	// cpu阻塞
	// 可以每一秒执行一次，防止干死机
	time.Sleep(time.Second * 1)
}
