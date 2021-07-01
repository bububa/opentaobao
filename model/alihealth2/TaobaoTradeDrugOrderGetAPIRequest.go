package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugOrderGetAPIRequest
查看订单详情 API请求
taobao.trade.drug.order.get

商家查看订单详情 */
type TaobaoTradeDrugOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// New
