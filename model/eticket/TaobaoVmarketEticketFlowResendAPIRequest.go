package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketFlowResendAPIRequest
业务重新触发发码短信 API请求
taobao.vmarket.eticket.flow.resend

业务重新触发发码短信 */
type TaobaoVmarketEticketFlowResendAPIRequest struct {
	model.Params
	// 业务单号
	_outerId string
	// 业务类型值，可联系淘宝业务运营取得具体值
	_bizType int64
}

// New
