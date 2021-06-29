package consul

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"math/rand"
)

/** 从 consul 中发现服务
根据服务名发现服务
*/
func AgentHealthServiceByName() (string, error) {
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
		return "", err
	}
	agent := c.Agent()
	state, outs, aerr := agent.AgentHealthServiceByName(consulRegisterThisName)
	if aerr != nil {
		fmt.Println("从 consul 发现服务错误", aerr)
		return "", aerr
	}
	fmt.Println(state)
	fmt.Println(outs)
	
	for _, i2 := range outs {
		fmt.Println("根据服务名发现服务：", i2, i2.Service)
	}
	
	// 随机返回一个服务
	index := 0
	len := len(outs)
	if len == 0 {
		return "", errors.New(consulRegisterThisName + "无可用的服务。")
	}
	if len > 1 {
		index = rand.Intn(len - 1)
	}
	service := outs[index].Service
	host := fmt.Sprintf("%s:%d", service.Address, service.Port)
	fmt.Println("随机返回一个发现的服务", host)
	
	fmt.Println(aerr)
	return host, nil
}
