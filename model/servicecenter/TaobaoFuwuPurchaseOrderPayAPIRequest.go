package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuPurchaseOrderPayAPIRequest
内购服务订单付款页获取接口 API请求
taobao.fuwu.purchase.order.pay

通过接口获取某一订单的付款页面链接 */
type TaobaoFuwuPurchaseOrderPayAPIRequest struct {
	model.Params
	// APPKEY，必填
	_appkey string
	// 订单号，与外部订单号二选一
	_orderId int64
	// 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
	_outOrderId string
	// 设备类型，目前只支持PC，可选
	_deviceType string
}

// New
