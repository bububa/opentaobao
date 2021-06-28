package aliyun

// Listener 
/* model for simplify = false
type Listener struct {

    // ListenerPort
    
    Listenerport   int64 `json:"listenerport,omitempty"`
    

    // BackendServers
    
    Backendservers  struct {
        BackendServerHealth  []BackendServerHealth `json:"backend_server_health,omitempty"`
    } `json:"backendservers,omitempty"`
    

}
*/

// Listener 
type Listener struct {

    // ListenerPort
    Listenerport   int64 `json:"listenerport,omitempty"`

    // BackendServers
    Backendservers   []BackendServerHealth `json:"backendservers,omitempty"`

}
