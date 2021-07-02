package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseQueryloanplans 天猫开新车租后查询还款计划
// tmall.car.lease.queryloanplans
//
// 天猫开新车租后查询还款计划
func TmallCarLeaseQueryloanplans(clt *core.SDKClient, req *tmallcar.TmallCarLeaseQueryloanplansAPIRequest, session string) (*tmallcar.TmallCarLeaseQueryloanplansAPIResponse, error) {
	var resp tmallcar.TmallCarLeaseQueryloanplansAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
