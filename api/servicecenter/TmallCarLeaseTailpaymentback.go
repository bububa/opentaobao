package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseTailpaymentback 尾款处置方案回传
// tmall.car.lease.tailpaymentback
//
// 尾款处置方案回传
func TmallCarLeaseTailpaymentback(clt *core.SDKClient, req *servicecenter.TmallCarLeaseTailpaymentbackAPIRequest, resp *servicecenter.TmallCarLeaseTailpaymentbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
