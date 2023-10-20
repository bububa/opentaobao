package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyBillVerificateCallback 翱象ERP核销回调
// alibaba.tcls.aelophy.bill.verificate.callback
//
// 翱象ERP核销回调
func AlibabaTclsAelophyBillVerificateCallback(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillVerificateCallbackAPIRequest, resp *wdk.AlibabaTclsAelophyBillVerificateCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
