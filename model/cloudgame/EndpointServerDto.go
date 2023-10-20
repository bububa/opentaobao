package cloudgame

import (
	"sync"
)

// EndpointServerDto 结构体
type EndpointServerDto struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// isp
	Isp string `json:"isp,omitempty" xml:"isp,omitempty"`
	// webSocketDomain
	WebSocketDomain string `json:"web_socket_domain,omitempty" xml:"web_socket_domain,omitempty"`
	// wsToken
	WsToken string `json:"ws_token,omitempty" xml:"ws_token,omitempty"`
	// webSocketPort
	WebSocketPort int64 `json:"web_socket_port,omitempty" xml:"web_socket_port,omitempty"`
}

var poolEndpointServerDto = sync.Pool{
	New: func() any {
		return new(EndpointServerDto)
	},
}

// GetEndpointServerDto() 从对象池中获取EndpointServerDto
func GetEndpointServerDto() *EndpointServerDto {
	return poolEndpointServerDto.Get().(*EndpointServerDto)
}

// ReleaseEndpointServerDto 释放EndpointServerDto
func ReleaseEndpointServerDto(v *EndpointServerDto) {
	v.Name = ""
	v.Type = ""
	v.Isp = ""
	v.WebSocketDomain = ""
	v.WsToken = ""
	v.WebSocketPort = 0
	poolEndpointServerDto.Put(v)
}
