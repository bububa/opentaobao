package baichuan

// ResultData 
type ResultData struct {

    // 商品数量
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

    // 商品id列表
    
    ItemList   []int64 `json:"item_list,omitempty" xml:"item_list>int64,omitempty"`
    

    // 具体内容
    
    DataList   []int64 `json:"data_list,omitempty" xml:"data_list>int64,omitempty"`
    

}
