package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTaSmsNumSend 短信发送
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
func AlibabaAliqinTaSmsNumSend(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinTaSmsNumSendAPIRequest, resp *alicom.AlibabaAliqinTaSmsNumSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
