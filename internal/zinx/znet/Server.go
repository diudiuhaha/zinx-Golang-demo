package znet

import (
	"fmt"
	"log"
	"net"
	"zinx/internal/zinx/ziface"
)

// Server 根据IServer的接口实现，定义一个Server的服务模块
type Server struct { // 定义四个属性
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IpVersion string
	// 服务器监听的IP
	Ip string
	// 服务器监听的端口
	Port int
}

// 建立实现接口的方法

// InitServer 实现初始化的方法
func InitServer(name string) ziface.IServer {
	// 创建一个对象,结构体的实例化
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8080,
	}
	return s
}

// Start 实现启动的方法
func (s *Server) Start() {
	log.Printf("[start] Server Listenner at IP %v:%v ,now is starting! \n", s.Ip, s.Port)
	// 1. 获取地址
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		log.Println("resolve tcp addr error :", err)
		return
	}
	// 2. 监听地址
	tcpListener, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		log.Printf("listen %v %v:", addr, err)
		return
	}
	log.Printf("start Zinx Server sucess,%v listenning at %v \n", s.Name, addr)
	// 3. 阻塞等待客户端连接,处理客户端连接连接（读写）
	for { // 建立一个循环，如果有客户端过来，会阻塞返回
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Println("accept err :", err)
			continue
		}
		// 当前已经与客户端建立连接
		// 可以建立一个空的业务，最大512字节的回显
		go func() {
			for {
				msg := make([]byte, 512)
				read, err := tcpConn.Read(msg) // 读取到客户端发的信息，取512字节
				if err != nil {
					log.Println("read msg err", err)
					continue
				}
				log.Printf("recv client buf is %s ,cnt is %d,", msg, read)
				// 回显功能
				// write, err := tcpConn.Write(msg[:read]) // 回写信息切片
				// if err != nil {
				// 	log.Println("write error :", write, err)
				// 	continue
				// }
				// 可以这样写
				if _, err := tcpConn.Write(msg[:read]); err != nil {
					log.Println("write err:", err)
				}

			}
		}()
	}
}

// Stop 实现停止的方法
func (s *Server) Stop() {
	// TODO 将服务器的资源/状态或者一些已经开辟的链接信息进行停止或者回收
}

// Serve 实现运行的方法
func (s *Server) Serve() {
	// 整个客户端处于阻塞的状态
	// 阻塞start()
	s.Start()
	// TODO 做一些启动服务之后的业务

	// 阻塞状态：
	select {}
}
