package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleasepayforcustomer 天猫开新车租后代客户还款
// tmall.car.lease.payforcustomer
//
// 天猫开新车租后代客户还款
func Tmallcarleasepayforcustomer(clt *core.SDKClient, req *tmallcar.TmallcarleasepayforcustomerAPIRequest, session string) (*tmallcar.TmallcarleasepayforcustomerAPIResponse, error) {
	var resp tmallcar.TmallcarleasepayforcustomerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
