package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtSmsCodeSend 喵零发送短信
// tmall.nrt.sms.code.send
//
// 喵零发送短信
func TmallNrtSmsCodeSend(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtSmsCodeSendAPIRequest, resp *nrt.TmallNrtSmsCodeSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
