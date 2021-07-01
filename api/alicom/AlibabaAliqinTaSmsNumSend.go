package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAliqinTaSmsNumSend
短信发送
alibaba.aliqin.ta.sms.num.send

短信发送 */
func AlibabaAliqinTaSmsNumSend(clt *core.SDKClient, req *alicom.AlibabaAliqinTaSmsNumSendAPIRequest, session string) (*alicom.AlibabaAliqinTaSmsNumSendAPIResponse, error) {
	var resp alicom.AlibabaAliqinTaSmsNumSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
