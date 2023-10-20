package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtSmsCodeSend 喵零发送短信
// tmall.nrt.sms.code.send
//
// 喵零发送短信
func TmallNrtSmsCodeSend(clt *core.SDKClient, req *nrt.TmallNrtSmsCodeSendAPIRequest, resp *nrt.TmallNrtSmsCodeSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
