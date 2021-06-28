package qimen

// BatchItemSynItem 
/* model for simplify = false
type BatchItemSynItem struct {

    // 没有同步成功的商品的编码
    
    ItemCode   string `json:"itemCode,omitempty"`
    

    // 出错信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// BatchItemSynItem 
type BatchItemSynItem struct {

    // 没有同步成功的商品的编码
    ItemCode   string `json:"itemCode,omitempty"`

    // 出错信息
    Message   string `json:"message,omitempty"`

}
