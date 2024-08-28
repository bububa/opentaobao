package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyOrderReceiptQuery 订单小票查询
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
func AlibabaTclsAelophyOrderReceiptQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyOrderReceiptQueryAPIRequest, resp *wdk.AlibabaTclsAelophyOrderReceiptQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
