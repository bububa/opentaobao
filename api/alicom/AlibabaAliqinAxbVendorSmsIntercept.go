package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinAxbVendorSmsIntercept AXB短信托收推送接口
// alibaba.aliqin.axb.vendor.sms.intercept
//
// 用于给供应商推送需要托收的短信
func AlibabaAliqinAxbVendorSmsIntercept(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorSmsInterceptAPIRequest, resp *alicom.AlibabaAliqinAxbVendorSmsInterceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
