package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseTailpaymentback 尾款处置方案回传
// tmall.car.lease.tailpaymentback
//
// 尾款处置方案回传
func TmallCarLeaseTailpaymentback(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallCarLeaseTailpaymentbackAPIRequest, resp *servicecenter.TmallCarLeaseTailpaymentbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
