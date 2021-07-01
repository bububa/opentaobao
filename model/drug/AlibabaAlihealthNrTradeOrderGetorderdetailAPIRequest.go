package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest
根据订单id获取单条订单详情 API请求
alibaba.alihealth.nr.trade.order.getorderdetail

阿里健康O2O，获取订单详情，修复组合商品价格精度问题 */
type AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// New
