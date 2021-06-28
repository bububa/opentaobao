package category

// WholesaleCategory 
type WholesaleCategory struct {

    // 类目路径
    
    Paths   []string `json:"paths,omitempty" xml:"paths>string,omitempty"`
    

    // 类目节点map
    
    NodeMap   string `json:"node_map,omitempty" xml:"node_map,omitempty"`
    

}
