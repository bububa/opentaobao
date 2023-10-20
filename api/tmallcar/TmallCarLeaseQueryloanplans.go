package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleasequeryloanplans 天猫开新车租后查询还款计划
// tmall.car.lease.queryloanplans
//
// 天猫开新车租后查询还款计划
func Tmallcarleasequeryloanplans(clt *core.SDKClient, req *tmallcar.TmallcarleasequeryloanplansAPIRequest, session string) (*tmallcar.TmallcarleasequeryloanplansAPIResponse, error) {
	var resp tmallcar.TmallcarleasequeryloanplansAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
