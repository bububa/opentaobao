package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionDeliveryCalculate 计算渠道用户下单快递费
// alibaba.damai.maitix.distribution.delivery.calculate
//
// 计算渠道用户下单快递费
func AlibabaDamaiMaitixDistributionDeliveryCalculate(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
