package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupMemberInfosAPIRequest
组团购批量获取用户参团信息 API请求
taobao.opentrade.group.member.infos

组团购场景下，获取用户参团信息 */
type TaobaoOpentradeGroupMemberInfosAPIRequest struct {
	model.Params
	// 团id
	_groupId int64
	// 用户openId列表
	_openUserIds []string
}

// New
