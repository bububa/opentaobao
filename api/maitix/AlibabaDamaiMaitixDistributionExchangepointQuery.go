package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixdistributionexchangepointquery 分销查询取票点接口
// alibaba.damai.maitix.distribution.exchangepoint.query
//
// 分销查询取票点接口
func Alibabadamaimaitixdistributionexchangepointquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixdistributionexchangepointqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixdistributionexchangepointqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixdistributionexchangepointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
