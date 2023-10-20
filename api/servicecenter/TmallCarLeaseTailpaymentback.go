package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarleasetailpaymentback 尾款处置方案回传
// tmall.car.lease.tailpaymentback
//
// 尾款处置方案回传
func Tmallcarleasetailpaymentback(clt *core.SDKClient, req *servicecenter.TmallcarleasetailpaymentbackAPIRequest, session string) (*servicecenter.TmallcarleasetailpaymentbackAPIResponse, error) {
	var resp servicecenter.TmallcarleasetailpaymentbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
