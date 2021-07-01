package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

/* TmallCarLeasePayforcustomer
天猫开新车租后代客户还款
tmall.car.lease.payforcustomer

天猫开新车租后代客户还款 */
func TmallCarLeasePayforcustomer(clt *core.SDKClient, req *tmallcar.TmallCarLeasePayforcustomerAPIRequest, session string) (*tmallcar.TmallCarLeasePayforcustomerAPIResponse, error) {
	var resp tmallcar.TmallCarLeasePayforcustomerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
