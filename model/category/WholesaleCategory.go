package category

// WholesaleCategory 
/* model for simplify = false
type WholesaleCategory struct {

    // 类目路径
    
    Paths  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"paths,omitempty"`
    

    // 类目节点map
    
    NodeMap   string `json:"node_map,omitempty"`
    

}
*/

// WholesaleCategory 
type WholesaleCategory struct {

    // 类目路径
    Paths   []string `json:"paths,omitempty"`

    // 类目节点map
    NodeMap   string `json:"node_map,omitempty"`

}
