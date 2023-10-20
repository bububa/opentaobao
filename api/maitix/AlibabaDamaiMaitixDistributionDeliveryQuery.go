package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixdistributiondeliveryquery 查询分销物流单
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
func Alibabadamaimaitixdistributiondeliveryquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixdistributiondeliveryqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixdistributiondeliveryqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixdistributiondeliveryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
