package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotSmsSend 物联网短信发送
// alibaba.aliqin.fc.iot.sms.send
//
// 发送物联网短信，只允许使用物联网短信模板
func AlibabaAliqinFcIotSmsSend(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotSmsSendAPIRequest, resp *aliqin.AlibabaAliqinFcIotSmsSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
