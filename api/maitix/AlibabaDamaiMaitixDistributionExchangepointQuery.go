package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionExchangepointQuery 分销查询取票点接口
// alibaba.damai.maitix.distribution.exchangepoint.query
//
// 分销查询取票点接口
func AlibabaDamaiMaitixDistributionExchangepointQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
