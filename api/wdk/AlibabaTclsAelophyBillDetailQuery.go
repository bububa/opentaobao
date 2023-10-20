package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyBillDetailQuery 账单明细接口
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
func AlibabaTclsAelophyBillDetailQuery(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillDetailQueryAPIRequest, resp *wdk.AlibabaTclsAelophyBillDetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
