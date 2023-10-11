package alihouse

import (
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
