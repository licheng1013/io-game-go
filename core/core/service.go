package core

import (
	"core/common"
	"core/register"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Service 构建服务
type Service struct {
	Nacos *register.Nacos
	Port  uint64
	Ip string
	rpcCell func(ip string,port uint64)
}

// NewService 设置启动端口地址
func NewService(ip string, port uint64) *Service {
	log.SetFlags(log.LstdFlags + log.Lshortfile)
	g := &Service{}
	g.Port = port
	g.Ip = ip
	g.Nacos = register.NewNacos()
	return g
}

// Run 设置注册中心地址和端口
func (n *Service) Run(ip string, port uint64) {
	log.Println("nacos注册地址: http://"+ ip +":"+fmt.Sprint(port))
	n.Nacos.SetServerConfig(n.Ip, port)
	n.Nacos.Register(ip, n.Port, common.ServicerName)
	n.Nacos.Init()      //初始化
	n.Nacos.Heartbeat() //心跳服务
	n.rpcCell(n.Ip,n.Port)        //注册rpc
	n.Stop()
	n.Nacos.Logout()
}

func (n *Service) RpcLient(f func(ip string,port uint64)) {
	n.rpcCell = f
}

func (n *Service) Stop() {
	log.Println("等待关闭...")
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("正在关机...")
}
