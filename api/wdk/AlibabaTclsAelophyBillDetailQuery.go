package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyBillDetailQuery 账单明细接口
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
func AlibabaTclsAelophyBillDetailQuery(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillDetailQueryAPIRequest, session string) (*wdk.AlibabaTclsAelophyBillDetailQueryAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyBillDetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
