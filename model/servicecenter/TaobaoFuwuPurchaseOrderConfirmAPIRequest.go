package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuPurchaseOrderConfirmAPIRequest
服务市场内购服务下单接口 API请求
taobao.fuwu.purchase.order.confirm

通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接 */
type TaobaoFuwuPurchaseOrderConfirmAPIRequest struct {
	model.Params
	// 内购服务下单接口参数
	_paramOrderConfirmQueryDTO *OrderConfirmQueryDto
}

// New
