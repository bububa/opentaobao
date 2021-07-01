package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupOpenAPIRequest
组团购场景开团 API请求
taobao.opentrade.group.open

组团购场景下，团长开团 */
type TaobaoOpentradeGroupOpenAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 用户openId
	_openUserId string
	// 组团活动id
	_groupActivityId int64
}

// New
