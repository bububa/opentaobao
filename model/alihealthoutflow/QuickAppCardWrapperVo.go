package alihealthoutflow

import (
	"sync"
)

// QuickAppCardWrapperVo 结构体
type QuickAppCardWrapperVo struct {
	// 疫苗卡片数据集合
	ItemInStockCardVoList []ItemInStockCardVo `json:"item_in_stock_card_vo_list,omitempty" xml:"item_in_stock_card_vo_list>item_in_stock_card_vo,omitempty"`
	// 投放卡片
	QuickAppCardInfoVoList []QuickAppCardInfoVo `json:"quick_app_card_info_vo_list,omitempty" xml:"quick_app_card_info_vo_list>quick_app_card_info_vo,omitempty"`
	// 地址信息
	AddressVo *AddressVo `json:"address_vo,omitempty" xml:"address_vo,omitempty"`
	// 卡片类型，1投放卡片，2疫苗卡片
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolQuickAppCardWrapperVo = sync.Pool{
	New: func() any {
		return new(QuickAppCardWrapperVo)
	},
}

// GetQuickAppCardWrapperVo() 从对象池中获取QuickAppCardWrapperVo
func GetQuickAppCardWrapperVo() *QuickAppCardWrapperVo {
	return poolQuickAppCardWrapperVo.Get().(*QuickAppCardWrapperVo)
}

// ReleaseQuickAppCardWrapperVo 释放QuickAppCardWrapperVo
func ReleaseQuickAppCardWrapperVo(v *QuickAppCardWrapperVo) {
	v.ItemInStockCardVoList = v.ItemInStockCardVoList[:0]
	v.QuickAppCardInfoVoList = v.QuickAppCardInfoVoList[:0]
	v.AddressVo = nil
	v.Type = 0
	poolQuickAppCardWrapperVo.Put(v)
}
