package ascpchannel

// Itemdto 
type Itemdto struct {
    // 货品ID
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 前端商品列表
    ItemDOList   []Itemdolist `json:"item_d_o_list,omitempty" xml:"item_d_o_list>itemdolist,omitempty"`
    // 前端商品信息
    ItemDoList   []Itemdolist `json:"item_do_list,omitempty" xml:"item_do_list>itemdolist,omitempty"`
}
