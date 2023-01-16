package ziface

// 定义一个接口服务器，实现三个方法
type IServer interface {
	// Start 启动
	Start()
	// Stop 停止
	Stop()
	// Serve 运行
	Serve()
}
