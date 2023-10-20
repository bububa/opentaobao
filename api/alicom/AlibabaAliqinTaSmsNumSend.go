package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTaSmsNumSend 短信发送
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
func AlibabaAliqinTaSmsNumSend(clt *core.SDKClient, req *alicom.AlibabaAliqinTaSmsNumSendAPIRequest, resp *alicom.AlibabaAliqinTaSmsNumSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
