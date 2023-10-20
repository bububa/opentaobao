package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinAxbVendorPushCallRelease 供应商推送通话结束事件
// alibaba.aliqin.axb.vendor.push.call.release
//
// 通话结束挂断的时候，供应商推送通话结束事件给阿里侧
func AlibabaAliqinAxbVendorPushCallRelease(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorPushCallReleaseAPIRequest, resp *alicom.AlibabaAliqinAxbVendorPushCallReleaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
