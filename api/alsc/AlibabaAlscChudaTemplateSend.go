package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscChudaTemplateSend 本地生活触达模板消息发送接口
// alibaba.alsc.chuda.template.send
//
// 允许三方小程序通过该api发送本地生活触达消息
func AlibabaAlscChudaTemplateSend(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscChudaTemplateSendAPIRequest, resp *alsc.AlibabaAlscChudaTemplateSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
