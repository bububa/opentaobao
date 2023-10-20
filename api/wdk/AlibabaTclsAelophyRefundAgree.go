package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundAgree saas 售后逆向 商户同意用户逆向申请
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
func AlibabaTclsAelophyRefundAgree(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundAgreeAPIRequest, resp *wdk.AlibabaTclsAelophyRefundAgreeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
