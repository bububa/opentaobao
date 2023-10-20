package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderDtdResend 门店自送重发码
// taobao.omniorder.dtd.resend
//
// 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
func TaobaoOmniorderDtdResend(clt *core.SDKClient, req *omniorder.TaobaoOmniorderDtdResendAPIRequest, resp *omniorder.TaobaoOmniorderDtdResendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
