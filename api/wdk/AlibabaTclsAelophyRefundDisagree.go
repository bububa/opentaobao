package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundDisagree saas 售后逆向 商户拒绝用户逆向申请
// alibaba.tcls.aelophy.refund.disagree
//
// saas 售后逆向 商户拒绝用户逆向申请
func AlibabaTclsAelophyRefundDisagree(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundDisagreeAPIRequest, resp *wdk.AlibabaTclsAelophyRefundDisagreeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
