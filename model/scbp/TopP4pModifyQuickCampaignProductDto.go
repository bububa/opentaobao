package scbp

import (
	"sync"
)

// TopP4pModifyQuickCampaignProductDto 结构体
type TopP4pModifyQuickCampaignProductDto struct {
	// 商品ID
	ProductIdList []int64 `json:"product_id_list,omitempty" xml:"product_id_list>int64,omitempty"`
	// 操作类型，0=商品暂停，1=商品开启，2=新增商品，3=删除商品
	Action int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolTopP4pModifyQuickCampaignProductDto = sync.Pool{
	New: func() any {
		return new(TopP4pModifyQuickCampaignProductDto)
	},
}

// GetTopP4pModifyQuickCampaignProductDto() 从对象池中获取TopP4pModifyQuickCampaignProductDto
func GetTopP4pModifyQuickCampaignProductDto() *TopP4pModifyQuickCampaignProductDto {
	return poolTopP4pModifyQuickCampaignProductDto.Get().(*TopP4pModifyQuickCampaignProductDto)
}

// ReleaseTopP4pModifyQuickCampaignProductDto 释放TopP4pModifyQuickCampaignProductDto
func ReleaseTopP4pModifyQuickCampaignProductDto(v *TopP4pModifyQuickCampaignProductDto) {
	v.ProductIdList = v.ProductIdList[:0]
	v.Action = 0
	v.CampaignId = 0
	poolTopP4pModifyQuickCampaignProductDto.Put(v)
}
