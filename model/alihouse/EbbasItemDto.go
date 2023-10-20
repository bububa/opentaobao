package alihouse

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// EbbasItemDto 结构体
type EbbasItemDto struct {
	// 服务者列表
	BrokerList []CommunityAgentRelationDto `json:"broker_list,omitempty" xml:"broker_list>community_agent_relation_dto,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 楼盘上下架状态
	OnlineStatus *model.File `json:"online_status,omitempty" xml:"online_status,omitempty"`
}

var poolEbbasItemDto = sync.Pool{
	New: func() any {
		return new(EbbasItemDto)
	},
}

// GetEbbasItemDto() 从对象池中获取EbbasItemDto
func GetEbbasItemDto() *EbbasItemDto {
	return poolEbbasItemDto.Get().(*EbbasItemDto)
}

// ReleaseEbbasItemDto 释放EbbasItemDto
func ReleaseEbbasItemDto(v *EbbasItemDto) {
	v.BrokerList = v.BrokerList[:0]
	v.ItemId = ""
	v.OuterId = ""
	v.OnlineStatus = nil
	poolEbbasItemDto.Put(v)
}
