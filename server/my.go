package server

import (
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"log"
)

//import "github.com/fatedier/frp/client"
type AllotPortReq struct {
	ProxyType  string `json:"type"`
	UsedType   uint32 `json:"used_type"` //用途：0=UsedClient=client,1=UsedSSH=ssh
	ServerUid  string `json:"server_uid"`
	RemoteName string `json:"remote_name"`
	RemoteIp   string `json:"remote_ip"`
	RemotePort int    `json:"remote_port"`
	Sign       string `json:"sign"`
}

func AddProxyByApi() {

	pxyCfgs := make(map[string]config.ProxyConf)
	//cfg := new(config.ProxyConf)
	common := config.GetDefaultClientConf()
	cfg := config.DefaultProxyConf("tcp")

	if tcpConfig, ok := cfg.(*config.TCPProxyConf); ok {
		tcpConfig.GetBaseInfo().ProxyName = "novnc"
		tcpConfig.GetBaseInfo().LocalIP = "127.0.0.1"
		tcpConfig.GetBaseInfo().LocalPort = 8866
		tcpConfig.RemotePort = 9090

		common.ServerAddr = "66.42.35.117"
		common.ServerPort = 7000

		pxyCfgs["novnc"] = tcpConfig
		svr, errRet := client.NewService(common, pxyCfgs, nil, "")
		if errRet != nil {
			log.Fatal(errRet)
		}

		svr.GetController()

		err := svr.Run()
		if err != nil {
			log.Fatal(err)
		}
		// ...use f...
	} else {
		log.Println("something wrong")
	}

	//prnm := proxy.NewManager()

	//content := &plugin.NewProxyContent{
	//	User: plugin.UserInfo{
	//		User:  ctl.loginMsg.User,
	//		Metas: ctl.loginMsg.Metas,
	//		RunID: ctl.loginMsg.RunID,
	//	},
	//	NewProxy: *m,
	//}
	//prnm.Add("xx", proxy.Proxy())
}
