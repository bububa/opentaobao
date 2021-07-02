package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseReserve 整车租车回传预约信息
// tmall.car.lease.reserve
//
// 租赁公司回传预约到店信息
func TmallCarLeaseReserve(clt *core.SDKClient, req *servicecenter.TmallCarLeaseReserveAPIRequest, session string) (*servicecenter.TmallCarLeaseReserveAPIResponse, error) {
	var resp servicecenter.TmallCarLeaseReserveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
