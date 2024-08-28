package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionDeliveryCalculate 计算渠道用户下单快递费
// alibaba.damai.maitix.distribution.delivery.calculate
//
// 计算渠道用户下单快递费
func AlibabaDamaiMaitixDistributionDeliveryCalculate(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
