package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundCsapplyNew 代客退
// alibaba.tcls.aelophy.refund.csapply.new
//
// 代客退
func AlibabaTclsAelophyRefundCsapplyNew(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundCsapplyNewAPIRequest, resp *wdk.AlibabaTclsAelophyRefundCsapplyNewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
