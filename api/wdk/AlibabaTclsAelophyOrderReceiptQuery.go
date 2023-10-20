package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyOrderReceiptQuery 订单小票查询
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
func AlibabaTclsAelophyOrderReceiptQuery(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyOrderReceiptQueryAPIRequest, session string) (*wdk.AlibabaTclsAelophyOrderReceiptQueryAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyOrderReceiptQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
