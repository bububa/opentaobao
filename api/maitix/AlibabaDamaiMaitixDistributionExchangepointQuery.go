package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

/* AlibabaDamaiMaitixDistributionExchangepointQuery
分销查询取票点接口
alibaba.damai.maitix.distribution.exchangepoint.query

分销查询取票点接口 */
func AlibabaDamaiMaitixDistributionExchangepointQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse, error) {
	var resp maitix.AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
