package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeasePayforcustomer 天猫开新车租后代客户还款
// tmall.car.lease.payforcustomer
//
// 天猫开新车租后代客户还款
func TmallCarLeasePayforcustomer(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeasePayforcustomerAPIRequest, resp *tmallcar.TmallCarLeasePayforcustomerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
