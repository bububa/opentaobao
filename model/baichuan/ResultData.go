package baichuan

// ResultData 
type ResultData struct {

    // 商品数量
    Count   int64 `json:"count,omitempty"`

    // 商品id列表
    ItemList   []Number `json:"item_list,omitempty"`

    // 具体内容
    DataList   []Number `json:"data_list,omitempty"`

}
