package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOrderQuery 大麦-查询分销单
// alibaba.damai.maitix.order.query
//
// 查询分销单
func AlibabaDamaiMaitixOrderQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
