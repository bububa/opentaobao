package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseRiskcallback 整车租赁风控模型回调
// tmall.car.lease.riskcallback
//
// 租赁公司回调风控结果
func TmallCarLeaseRiskcallback(clt *core.SDKClient, req *servicecenter.TmallCarLeaseRiskcallbackAPIRequest, resp *servicecenter.TmallCarLeaseRiskcallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
