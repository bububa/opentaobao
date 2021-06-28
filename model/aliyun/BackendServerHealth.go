package aliyun

// BackendServerHealth 
/* model for simplify = false
type BackendServerHealth struct {

    // serverId
    
    Serverid   string `json:"serverid,omitempty"`
    

    // ServerHealthStatus<br/>后端服务器的健康状况，normal,abnormal或Unavailable。
    
    Serverhealthstatus   string `json:"serverhealthstatus,omitempty"`
    

}
*/

// BackendServerHealth 
type BackendServerHealth struct {

    // serverId
    Serverid   string `json:"serverid,omitempty"`

    // ServerHealthStatus<br/>后端服务器的健康状况，normal,abnormal或Unavailable。
    Serverhealthstatus   string `json:"serverhealthstatus,omitempty"`

}
