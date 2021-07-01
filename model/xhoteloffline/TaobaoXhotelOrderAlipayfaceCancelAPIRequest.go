package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceCancelAPIRequest
线下信用住订单取消 API请求
taobao.xhotel.order.alipayface.cancel

提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消 */
type TaobaoXhotelOrderAlipayfaceCancelAPIRequest struct {
	model.Params
	// 淘宝订单ID，必选
	_tid int64
	// 原因描述
	_reasonText string
	// 外部订单号
	_outId string
	// 预留后续用
	_notifyUrl string
	// 请求流水号
	_outUuid string
}

// New
