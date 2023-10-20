package alihouse

import (
	"sync"
)

// ProjectOrderDto 结构体
type ProjectOrderDto struct {
	// 交易商品列表
	TradeItemList []ProjectTradeItemOrderDto `json:"trade_item_list,omitempty" xml:"trade_item_list>project_trade_item_order_dto,omitempty"`
	// 模块排序信息
	ModuleOrder []ModuleTypeOrderDto `json:"module_order,omitempty" xml:"module_order>module_type_order_dto,omitempty"`
	// 楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
}

var poolProjectOrderDto = sync.Pool{
	New: func() any {
		return new(ProjectOrderDto)
	},
}

// GetProjectOrderDto() 从对象池中获取ProjectOrderDto
func GetProjectOrderDto() *ProjectOrderDto {
	return poolProjectOrderDto.Get().(*ProjectOrderDto)
}

// ReleaseProjectOrderDto 释放ProjectOrderDto
func ReleaseProjectOrderDto(v *ProjectOrderDto) {
	v.TradeItemList = v.TradeItemList[:0]
	v.ModuleOrder = v.ModuleOrder[:0]
	v.OuterId = ""
	v.OrderNo = 0
	poolProjectOrderDto.Put(v)
}
