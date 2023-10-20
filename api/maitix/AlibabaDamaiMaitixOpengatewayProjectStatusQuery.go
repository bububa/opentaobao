package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOpengatewayProjectStatusQuery 分销状态查询接口queryProjectStatusByProjectId
// alibaba.damai.maitix.opengateway.project.status.query
//
// queryProjectStatusByProjectId
func AlibabaDamaiMaitixOpengatewayProjectStatusQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
