package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFcVoiceNumCancelcall 取消呼叫
// alibaba.aliqin.fc.voice.num.cancelcall
//
// 当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
func AlibabaAliqinFcVoiceNumCancelcall(clt *core.SDKClient, req *alicom.AlibabaAliqinFcVoiceNumCancelcallAPIRequest, resp *alicom.AlibabaAliqinFcVoiceNumCancelcallAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
