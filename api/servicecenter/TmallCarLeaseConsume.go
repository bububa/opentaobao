package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseConsume 汽车租赁核销
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
func TmallCarLeaseConsume(clt *core.SDKClient, req *servicecenter.TmallCarLeaseConsumeAPIRequest, resp *servicecenter.TmallCarLeaseConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
