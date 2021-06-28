package fenxiao

// InventoryQueryResult 
/* model for simplify = false
type InventoryQueryResult struct {

    // 查询成功列表
    
    ItemInventorys  struct {
        InventoryInfoDetailDto  []InventoryInfoDetailDto `json:"inventory_info_detail_dto,omitempty"`
    } `json:"item_inventorys,omitempty"`
    

    // tipInfos
    
    TipInfos  struct {
        TipInfo  []TipInfo `json:"tip_info,omitempty"`
    } `json:"tip_infos,omitempty"`
    

}
*/

// InventoryQueryResult 
type InventoryQueryResult struct {

    // 查询成功列表
    ItemInventorys   []InventoryInfoDetailDto `json:"item_inventorys,omitempty"`

    // tipInfos
    TipInfos   []TipInfo `json:"tip_infos,omitempty"`

}
