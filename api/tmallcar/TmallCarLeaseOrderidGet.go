package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseOrderidGet 天猫开新车查询订单id
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
func TmallCarLeaseOrderidGet(clt *core.SDKClient, req *tmallcar.TmallCarLeaseOrderidGetAPIRequest, resp *tmallcar.TmallCarLeaseOrderidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
