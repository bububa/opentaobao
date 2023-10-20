package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixSeatTokenQuery 分销商选座获取qtoken
// alibaba.damai.maitix.seat.token.query
//
// 选座分销，分销商查询token
func AlibabaDamaiMaitixSeatTokenQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixSeatTokenQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixSeatTokenQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
