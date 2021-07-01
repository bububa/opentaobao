package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcVoiceNumCancelcallAPIRequest
取消呼叫 API请求
alibaba.aliqin.fc.voice.num.cancelcall

当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话 */
type AlibabaAliqinFcVoiceNumCancelcallAPIRequest struct {
	model.Params
	// 呼叫唯一id
	_callId string
}

// New
