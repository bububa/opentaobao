package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinAxbVendorSmsIntercept AXB短信托收推送接口
// alibaba.aliqin.axb.vendor.sms.intercept
//
// 用于给供应商推送需要托收的短信
func AlibabaAliqinAxbVendorSmsIntercept(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorSmsInterceptAPIRequest, resp *alicom.AlibabaAliqinAxbVendorSmsInterceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
