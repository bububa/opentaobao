package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotsmssend 物联网短信发送
// alibaba.aliqin.fc.iot.sms.send
//
// 发送物联网短信，只允许使用物联网短信模板
func Alibabaaliqinfciotsmssend(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotsmssendAPIRequest, session string) (*aliqin.AlibabaaliqinfciotsmssendAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotsmssendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
