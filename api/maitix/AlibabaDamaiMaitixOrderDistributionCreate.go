package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOrderDistributionCreate 大麦-新分销下单
// alibaba.damai.maitix.order.distribution.create
//
// createDistributionOrder
func AlibabaDamaiMaitixOrderDistributionCreate(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderDistributionCreateAPIRequest, resp *maitix.AlibabaDamaiMaitixOrderDistributionCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
