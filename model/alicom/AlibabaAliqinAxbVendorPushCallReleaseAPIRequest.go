package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorPushCallReleaseAPIRequest
供应商推送通话结束事件 API请求
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧 */
type AlibabaAliqinAxbVendorPushCallReleaseAPIRequest struct {
	model.Params
	// end_call_request
	_endCallRequest *EndCallRequest
}

// New
