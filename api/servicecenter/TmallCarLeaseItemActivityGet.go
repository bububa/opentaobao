package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarleaseitemactivityget 查询汽车租赁活动信息
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
func Tmallcarleaseitemactivityget(clt *core.SDKClient, req *servicecenter.TmallcarleaseitemactivitygetAPIRequest, session string) (*servicecenter.TmallcarleaseitemactivitygetAPIResponse, error) {
	var resp servicecenter.TmallcarleaseitemactivitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
