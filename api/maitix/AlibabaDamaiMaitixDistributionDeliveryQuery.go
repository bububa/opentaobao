package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionDeliveryQuery 查询分销物流单
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
func AlibabaDamaiMaitixDistributionDeliveryQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
