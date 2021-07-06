package aliyun

// Listener 结构体
type Listener struct {
	// BackendServers
	Backendservers []BackendServerHealth `json:"backendservers,omitempty" xml:"backendservers>backend_server_health,omitempty"`
	// ListenerPort
	Listenerport int64 `json:"listenerport,omitempty" xml:"listenerport,omitempty"`
}
