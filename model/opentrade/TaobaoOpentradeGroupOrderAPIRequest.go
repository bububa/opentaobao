package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupOrderAPIRequest
组团购获取订单列表 API请求
taobao.opentrade.group.order

组团购场景下，获取开团的订单列表 */
type TaobaoOpentradeGroupOrderAPIRequest struct {
	model.Params
	// 团id
	_groupId int64
}

// New
