package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupMemberInfoAPIRequest
组团购获取用户参团信息 API请求
taobao.opentrade.group.member.info

组团购场景下，获取用户参团信息 */
type TaobaoOpentradeGroupMemberInfoAPIRequest struct {
	model.Params
	// 团id
	_groupId int64
	// 用户openId
	_openUserId string
}

// New
