package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundCsapply 商家代客售后提交逆向申请
// alibaba.tcls.aelophy.refund.csapply
//
// 商家代客售后提交逆向申请
func AlibabaTclsAelophyRefundCsapply(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundCsapplyAPIRequest, session string) (*wdk.AlibabaTclsAelophyRefundCsapplyAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyRefundCsapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
