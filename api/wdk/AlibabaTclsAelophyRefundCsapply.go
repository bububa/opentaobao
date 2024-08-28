package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundCsapply 商家代客售后提交逆向申请
// alibaba.tcls.aelophy.refund.csapply
//
// 商家代客售后提交逆向申请
func AlibabaTclsAelophyRefundCsapply(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundCsapplyAPIRequest, resp *wdk.AlibabaTclsAelophyRefundCsapplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
