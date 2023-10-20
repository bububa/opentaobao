package alihouse

import (
	"sync"
)

// ProjectTradeItemOrderDto 结构体
type ProjectTradeItemOrderDto struct {
	// 关联的商品列表
	RelationTradeItem []ProjectTradeItemOrderDto `json:"relation_trade_item,omitempty" xml:"relation_trade_item>project_trade_item_order_dto,omitempty"`
	// 交易工具图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 交易工具跳转地址
	JumpValue string `json:"jump_value,omitempty" xml:"jump_value,omitempty"`
	// 模块ID
	ModuleId string `json:"module_id,omitempty" xml:"module_id,omitempty"`
	// 交易商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 类型
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolProjectTradeItemOrderDto = sync.Pool{
	New: func() any {
		return new(ProjectTradeItemOrderDto)
	},
}

// GetProjectTradeItemOrderDto() 从对象池中获取ProjectTradeItemOrderDto
func GetProjectTradeItemOrderDto() *ProjectTradeItemOrderDto {
	return poolProjectTradeItemOrderDto.Get().(*ProjectTradeItemOrderDto)
}

// ReleaseProjectTradeItemOrderDto 释放ProjectTradeItemOrderDto
func ReleaseProjectTradeItemOrderDto(v *ProjectTradeItemOrderDto) {
	v.RelationTradeItem = v.RelationTradeItem[:0]
	v.ImageUrl = ""
	v.JumpValue = ""
	v.ModuleId = ""
	v.ItemId = 0
	v.OrderNo = 0
	v.ItemType = 0
	poolProjectTradeItemOrderDto.Put(v)
}
