package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupJoinAPIRequest
组团购场景参团 API请求
taobao.opentrade.group.join

组团购场景下，用户参团 */
type TaobaoOpentradeGroupJoinAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 团id
	_groupId int64
	// 用户openId
	_openUserId string
}

// New
