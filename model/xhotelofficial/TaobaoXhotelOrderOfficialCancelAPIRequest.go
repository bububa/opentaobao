package xhotelofficial

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderOfficialCancelAPIRequest
官网信用住订单取消 API请求
taobao.xhotel.order.official.cancel

官网信用住订单取消 */
type TaobaoXhotelOrderOfficialCancelAPIRequest struct {
	model.Params
	// 淘宝订单号,必选
	_tid int64
	// 原因描述
	_reasonText string
	// 外部订单号
	_outId string
	// 请求流水号（必须传入）
	_outUuid string
	// 暂无意义，无需传入
	_notifyUrl string
}

// New
