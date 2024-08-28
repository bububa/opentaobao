package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundDisagree saas 售后逆向 商户拒绝用户逆向申请
// alibaba.tcls.aelophy.refund.disagree
//
// saas 售后逆向 商户拒绝用户逆向申请
func AlibabaTclsAelophyRefundDisagree(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundDisagreeAPIRequest, resp *wdk.AlibabaTclsAelophyRefundDisagreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
