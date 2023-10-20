package aliyun

import (
	"sync"
)

// Listener 结构体
type Listener struct {
	// BackendServers
	Backendservers []BackendServerHealth `json:"backendservers,omitempty" xml:"backendservers>backend_server_health,omitempty"`
	// ListenerPort
	Listenerport int64 `json:"listenerport,omitempty" xml:"listenerport,omitempty"`
}

var poolListener = sync.Pool{
	New: func() any {
		return new(Listener)
	},
}

// GetListener() 从对象池中获取Listener
func GetListener() *Listener {
	return poolListener.Get().(*Listener)
}

// ReleaseListener 释放Listener
func ReleaseListener(v *Listener) {
	v.Backendservers = v.Backendservers[:0]
	v.Listenerport = 0
	poolListener.Put(v)
}
