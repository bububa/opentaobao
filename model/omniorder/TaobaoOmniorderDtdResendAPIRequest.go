package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderDtdResendAPIRequest
门店自送重发码 API请求
taobao.omniorder.dtd.resend

该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次 */
type TaobaoOmniorderDtdResendAPIRequest struct {
	model.Params
	// 淘宝主订单ID
	_mainOrderId int64
}

// New
