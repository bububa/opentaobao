package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest
信用住订单取消结算接口 API请求
taobao.xhotel.order.alipayface.cancelsettle

信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。 */
type TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest struct {
	model.Params
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填
	_outId string
}

// New
