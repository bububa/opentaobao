package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseStatussynchronize 天猫开新车租后状态同步
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
func TmallCarLeaseStatussynchronize(clt *core.SDKClient, req *tmallcar.TmallCarLeaseStatussynchronizeAPIRequest, resp *tmallcar.TmallCarLeaseStatussynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
