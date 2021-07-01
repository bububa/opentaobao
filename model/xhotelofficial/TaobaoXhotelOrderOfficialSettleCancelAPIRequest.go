package xhotelofficial

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderOfficialSettleCancelAPIRequest
官网信用住取消结账 API请求
taobao.xhotel.order.official.settle.cancel

用于官网信用住取消结账 */
type TaobaoXhotelOrderOfficialSettleCancelAPIRequest struct {
	model.Params
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填
	_outId string
	// 暂时无意义，无需传入
	_notifyUrl string
	// 请求流水号
	_outUuid string
}

// New
