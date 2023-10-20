package alihouse

import (
	"sync"
)

// SyncHouseTradeItemDto 结构体
type SyncHouseTradeItemDto struct {
	// 1
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 1
	BillPayItemId int64 `json:"bill_pay_item_id,omitempty" xml:"bill_pay_item_id,omitempty"`
	// 1
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 1
	RelationItemBizType int64 `json:"relation_item_biz_type,omitempty" xml:"relation_item_biz_type,omitempty"`
	// 是否需要异步
	IsAsync int64 `json:"is_async,omitempty" xml:"is_async,omitempty"`
}

var poolSyncHouseTradeItemDto = sync.Pool{
	New: func() any {
		return new(SyncHouseTradeItemDto)
	},
}

// GetSyncHouseTradeItemDto() 从对象池中获取SyncHouseTradeItemDto
func GetSyncHouseTradeItemDto() *SyncHouseTradeItemDto {
	return poolSyncHouseTradeItemDto.Get().(*SyncHouseTradeItemDto)
}

// ReleaseSyncHouseTradeItemDto 释放SyncHouseTradeItemDto
func ReleaseSyncHouseTradeItemDto(v *SyncHouseTradeItemDto) {
	v.CommunityOuterId = ""
	v.OuterId = ""
	v.OuterStoreId = ""
	v.BusinessType = 0
	v.BillPayItemId = 0
	v.MerchantOpenId = 0
	v.RelationItemBizType = 0
	v.IsAsync = 0
	poolSyncHouseTradeItemDto.Put(v)
}
