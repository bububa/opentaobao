package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseQueryloanplans 天猫开新车租后查询还款计划
// tmall.car.lease.queryloanplans
//
// 天猫开新车租后查询还款计划
func TmallCarLeaseQueryloanplans(clt *core.SDKClient, req *tmallcar.TmallCarLeaseQueryloanplansAPIRequest, resp *tmallcar.TmallCarLeaseQueryloanplansAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
