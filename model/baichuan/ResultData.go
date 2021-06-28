package baichuan

// ResultData 
/* model for simplify = false
type ResultData struct {

    // 商品数量
    
    Count   int64 `json:"count,omitempty"`
    

    // 商品id列表
    
    ItemList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"item_list,omitempty"`
    

    // 具体内容
    
    DataList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"data_list,omitempty"`
    

}
*/

// ResultData 
type ResultData struct {

    // 商品数量
    Count   int64 `json:"count,omitempty"`

    // 商品id列表
    ItemList   []int64 `json:"item_list,omitempty"`

    // 具体内容
    DataList   []int64 `json:"data_list,omitempty"`

}
