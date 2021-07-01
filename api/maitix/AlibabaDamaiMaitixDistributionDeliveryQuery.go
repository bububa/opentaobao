package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

/* AlibabaDamaiMaitixDistributionDeliveryQuery
查询分销物流单
alibaba.damai.maitix.distribution.delivery.query

渠道查询物流订单 */
func AlibabaDamaiMaitixDistributionDeliveryQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse, error) {
	var resp maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
