package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	consulapi "github.com/hashicorp/consul/api"
)

const (
	// 运行官网 consul 服务主机
	consulHost string = "http://127.0.0.1:8500"
	// 运行自定义 consul 集群
	// consulHost string = "http://127.0.0.1:8510"
	
	// 注册服务的路径
	consulRegisterPath string = "/v1/agent/service/register"
	// 服务ID
	consulRegisterThisId      string = "go-tcp-service-006"
	// 服务名称
	consulRegisterThisName    string = "go-tcp-service"
	// 服务地址
	consulRegisterThisAddress string = "192.168.0.103" // consul 运行在容器中，需要发现该服务，所以需要宿主机地址
	// 服务端口
	consulRegisterThisPort    int    = 8200
)

type Server struct {
	ID      string
	Name    string
	Address string
	Port    int
}

func init() {
	// 注册服务：将此服务注册到 consul 中
	ServiceRegister() // 官方 api
	// ServiceRegisterByCustomer() // 自定义注册服务方法
}

/** consul 官方 api 注册服务方法
注册服务：将此服务注册到 consul 中
 */
func ServiceRegister() error {
	consulConfig := &consulapi.Config{
		Address:    consulHost,
		Scheme:     "",
		Datacenter: "",
		Transport:  nil,
		HttpClient: nil,
		HttpAuth:   nil,
		WaitTime:   0,
		Token:      "",
		TokenFile:  "",
		Namespace:  "",
		TLSConfig:  consulapi.TLSConfig{},
	}
	c, err := consulapi.NewClient(consulConfig)
	if err != nil {
		fmt.Println("初始化 consul 客户端错误", err)
		return err
	}
	reg := &consulapi.AgentServiceRegistration{
		Kind:              "",
		ID:                consulRegisterThisId,
		Name:              consulRegisterThisName,
		Tags:              nil,
		Port:              consulRegisterThisPort,
		Address:           consulRegisterThisAddress,
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta:              nil,
		Weights:           nil,
		Check:             nil,
		Checks:            nil,
		Proxy:             nil,
		Connect:           nil,
		Namespace:         "",
	}
	agent := c.Agent()
	aerr := agent.ServiceRegister(reg)
	if aerr != nil {
		fmt.Println("注册服务到 consul 错误", aerr)
		return aerr
	}
	
	return nil
}

/** 自定义注册服务方法
注册服务：将此服务注册到 consul 中
*/
func ServiceRegisterByCustomer() {
	server := &Server{
		ID:      consulRegisterThisId,
		Name:    consulRegisterThisName,
		Address: consulRegisterThisAddress,
		Port:    consulRegisterThisPort,
	}
	s, _ := json.Marshal(server) // bytes
	url := fmt.Sprintf("%s%s", consulHost, consulRegisterPath)
	req, _ := http.NewRequest("PUT", url, bytes.NewReader(s))
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	fmt.Println(res)
}

