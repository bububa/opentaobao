package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaTclsAelophyRefundDisagree
saas 售后逆向 商户拒绝用户逆向申请
alibaba.tcls.aelophy.refund.disagree

saas 售后逆向 商户拒绝用户逆向申请 */
func AlibabaTclsAelophyRefundDisagree(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundDisagreeAPIRequest, session string) (*wdk.AlibabaTclsAelophyRefundDisagreeAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyRefundDisagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
