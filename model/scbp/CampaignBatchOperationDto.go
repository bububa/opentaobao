package scbp

import (
	"sync"
)

// CampaignBatchOperationDto 结构体
type CampaignBatchOperationDto struct {
	// 具体操作实例
	CampaignOperationList []Campaignoperationlist `json:"campaign_operation_list,omitempty" xml:"campaign_operation_list>campaignoperationlist,omitempty"`
}

var poolCampaignBatchOperationDto = sync.Pool{
	New: func() any {
		return new(CampaignBatchOperationDto)
	},
}

// GetCampaignBatchOperationDto() 从对象池中获取CampaignBatchOperationDto
func GetCampaignBatchOperationDto() *CampaignBatchOperationDto {
	return poolCampaignBatchOperationDto.Get().(*CampaignBatchOperationDto)
}

// ReleaseCampaignBatchOperationDto 释放CampaignBatchOperationDto
func ReleaseCampaignBatchOperationDto(v *CampaignBatchOperationDto) {
	v.CampaignOperationList = v.CampaignOperationList[:0]
	poolCampaignBatchOperationDto.Put(v)
}
