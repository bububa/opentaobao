package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOpengatewayPerformStatusQuery 分销状态查询接口queryPerformStatusByPerformId
// alibaba.damai.maitix.opengateway.perform.status.query
//
// queryPerformStatusByPerformId
func AlibabaDamaiMaitixOpengatewayPerformStatusQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
