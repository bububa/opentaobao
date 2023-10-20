package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseItemActivityGet 查询汽车租赁活动信息
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
func TmallCarLeaseItemActivityGet(clt *core.SDKClient, req *servicecenter.TmallCarLeaseItemActivityGetAPIRequest, session string) (*servicecenter.TmallCarLeaseItemActivityGetAPIResponse, error) {
	var resp servicecenter.TmallCarLeaseItemActivityGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
