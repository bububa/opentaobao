package aliyun

import (
	"sync"
)

// BackendServerHealth 结构体
type BackendServerHealth struct {
	// serverId
	Serverid string `json:"serverid,omitempty" xml:"serverid,omitempty"`
	// ServerHealthStatus&lt;br/&gt;后端服务器的健康状况，normal,abnormal或Unavailable。
	Serverhealthstatus string `json:"serverhealthstatus,omitempty" xml:"serverhealthstatus,omitempty"`
}

var poolBackendServerHealth = sync.Pool{
	New: func() any {
		return new(BackendServerHealth)
	},
}

// GetBackendServerHealth() 从对象池中获取BackendServerHealth
func GetBackendServerHealth() *BackendServerHealth {
	return poolBackendServerHealth.Get().(*BackendServerHealth)
}

// ReleaseBackendServerHealth 释放BackendServerHealth
func ReleaseBackendServerHealth(v *BackendServerHealth) {
	v.Serverid = ""
	v.Serverhealthstatus = ""
	poolBackendServerHealth.Put(v)
}
