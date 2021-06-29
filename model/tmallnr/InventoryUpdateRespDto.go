package tmallnr

// InventoryUpdateRespDTO 
type InventoryUpdateRespDTO struct {
    // 返回成功的库存记录数
    SuccessInventorys   []NrInventoryCheckDetailDTO `json:"success_inventorys,omitempty" xml:"success_inventorys>nr_inventory_check_detail_dto,omitempty"`
    // 失败的库存更新记录
    FailInventorys   []NrInventoryCheckDetailDTO `json:"fail_inventorys,omitempty" xml:"fail_inventorys>nr_inventory_check_detail_dto,omitempty"`
}
