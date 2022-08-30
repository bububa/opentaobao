package cloudgame

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
