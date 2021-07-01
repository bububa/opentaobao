package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderOfflineSettleCancelAPIRequest
线下信用住取消结账专用接口 API请求
taobao.xhotel.order.offline.settle.cancel

线下信用住取消结账专用接口 */
type TaobaoXhotelOrderOfflineSettleCancelAPIRequest struct {
	model.Params
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填（建议都写入）
	_outId string
	// 暂时无意义，无需传入
	_notifyUrl string
	// 请求流水号
	_outUuid string
}

// New
