package scbp

import (
	"sync"
)

// TopP4pModifyQuickCampaignTagDto 结构体
type TopP4pModifyQuickCampaignTagDto struct {
	// 定向标签编码，新增/修改有效
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 操作类型，1=修改定向溢价，2=创建定向标签，3=删除定向标签
	Action int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 定向溢价比例，(100,400]，新增/修改有效
	BidRate int64 `json:"bid_rate,omitempty" xml:"bid_rate,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 定向ID，删除有效
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 定向类型，新增/修改有效，1=类目访客，2=类目询盘，3=固定标签，比如本店、高MOQ等，4=地域定向
	TagType int64 `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
}

var poolTopP4pModifyQuickCampaignTagDto = sync.Pool{
	New: func() any {
		return new(TopP4pModifyQuickCampaignTagDto)
	},
}

// GetTopP4pModifyQuickCampaignTagDto() 从对象池中获取TopP4pModifyQuickCampaignTagDto
func GetTopP4pModifyQuickCampaignTagDto() *TopP4pModifyQuickCampaignTagDto {
	return poolTopP4pModifyQuickCampaignTagDto.Get().(*TopP4pModifyQuickCampaignTagDto)
}

// ReleaseTopP4pModifyQuickCampaignTagDto 释放TopP4pModifyQuickCampaignTagDto
func ReleaseTopP4pModifyQuickCampaignTagDto(v *TopP4pModifyQuickCampaignTagDto) {
	v.OptionValue = ""
	v.Action = 0
	v.BidRate = 0
	v.CampaignId = 0
	v.TagId = 0
	v.TagType = 0
	poolTopP4pModifyQuickCampaignTagDto.Put(v)
}
