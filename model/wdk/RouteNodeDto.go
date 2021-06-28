package wdk

// RouteNodeDto 
/* model for simplify = false
type RouteNodeDto struct {

    // 节点序号
    
    NodeIndex   int64 `json:"node_index,omitempty"`
    

    // 节点类型
    
    NodeType   string `json:"node_type,omitempty"`
    

    // 节点名称
    
    NodeCode   string `json:"node_code,omitempty"`
    

}
*/

// RouteNodeDto 
type RouteNodeDto struct {

    // 节点序号
    NodeIndex   int64 `json:"node_index,omitempty"`

    // 节点类型
    NodeType   string `json:"node_type,omitempty"`

    // 节点名称
    NodeCode   string `json:"node_code,omitempty"`

}
