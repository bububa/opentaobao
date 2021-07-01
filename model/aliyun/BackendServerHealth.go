package aliyun

// BackendServerHealth 结构体
type BackendServerHealth struct {
	// serverId
	Serverid string `json:"serverid,omitempty" xml:"serverid,omitempty"`
	// ServerHealthStatus<br/>后端服务器的健康状况，normal,abnormal或Unavailable。
	Serverhealthstatus string `json:"serverhealthstatus,omitempty" xml:"serverhealthstatus,omitempty"`
}
