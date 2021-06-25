package aliyun

// Listener 
type Listener struct {

    // ListenerPort
    Listenerport   int64 `json:"listenerport,omitempty"`

    // BackendServers
    Backendservers   []BackendServerHealth `json:"backendservers,omitempty"`

}
