package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseRiskcallback 整车租赁风控模型回调
// tmall.car.lease.riskcallback
//
// 租赁公司回调风控结果
func TmallCarLeaseRiskcallback(clt *core.SDKClient, req *servicecenter.TmallCarLeaseRiskcallbackAPIRequest, session string) (*servicecenter.TmallCarLeaseRiskcallbackAPIResponse, error) {
	var resp servicecenter.TmallCarLeaseRiskcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
