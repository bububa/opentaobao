package fenxiao

// InventoryQueryResult 
type InventoryQueryResult struct {
    // 查询成功列表
    ItemInventorys   []InventoryInfoDetailDto `json:"item_inventorys,omitempty" xml:"item_inventorys>inventory_info_detail_dto,omitempty"`
    // tipInfos
    TipInfos   []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
}
