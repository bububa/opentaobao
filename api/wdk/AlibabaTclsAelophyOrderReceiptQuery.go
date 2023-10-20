package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyOrderReceiptQuery 订单小票查询
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
func AlibabaTclsAelophyOrderReceiptQuery(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyOrderReceiptQueryAPIRequest, resp *wdk.AlibabaTclsAelophyOrderReceiptQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
