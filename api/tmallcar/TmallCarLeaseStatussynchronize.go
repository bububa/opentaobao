package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseStatussynchronize 天猫开新车租后状态同步
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
func TmallCarLeaseStatussynchronize(clt *core.SDKClient, req *tmallcar.TmallCarLeaseStatussynchronizeAPIRequest, session string) (*tmallcar.TmallCarLeaseStatussynchronizeAPIResponse, error) {
	var resp tmallcar.TmallCarLeaseStatussynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
