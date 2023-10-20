package tmallnr

import (
	"sync"
)

// InventoryUpdateRespDto 结构体
type InventoryUpdateRespDto struct {
	// 返回成功的库存记录数
	SuccessInventorys []NrInventoryCheckDetailDto `json:"success_inventorys,omitempty" xml:"success_inventorys>nr_inventory_check_detail_dto,omitempty"`
	// 失败的库存更新记录
	FailInventorys []NrInventoryCheckDetailDto `json:"fail_inventorys,omitempty" xml:"fail_inventorys>nr_inventory_check_detail_dto,omitempty"`
}

var poolInventoryUpdateRespDto = sync.Pool{
	New: func() any {
		return new(InventoryUpdateRespDto)
	},
}

// GetInventoryUpdateRespDto() 从对象池中获取InventoryUpdateRespDto
func GetInventoryUpdateRespDto() *InventoryUpdateRespDto {
	return poolInventoryUpdateRespDto.Get().(*InventoryUpdateRespDto)
}

// ReleaseInventoryUpdateRespDto 释放InventoryUpdateRespDto
func ReleaseInventoryUpdateRespDto(v *InventoryUpdateRespDto) {
	v.SuccessInventorys = v.SuccessInventorys[:0]
	v.FailInventorys = v.FailInventorys[:0]
	poolInventoryUpdateRespDto.Put(v)
}
