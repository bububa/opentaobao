package waybill

// Item 
/* model for simplify = false
type Item struct {

    // 数量
    
    Count   int64 `json:"count,omitempty"`
    

    // 名称
    
    Name   string `json:"name,omitempty"`
    

}
*/

// Item 
type Item struct {

    // 数量
    Count   int64 `json:"count,omitempty"`

    // 名称
    Name   string `json:"name,omitempty"`

}
