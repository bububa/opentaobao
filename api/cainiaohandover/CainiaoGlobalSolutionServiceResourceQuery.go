package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalSolutionServiceResourceQuery 查询解决方案服务资源列表
// cainiao.global.solution.service.resource.query
//
// 返回直接解决方案的指定物流服务的可用资源列表
func CainiaoGlobalSolutionServiceResourceQuery(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalSolutionServiceResourceQueryAPIRequest, resp *cainiaohandover.CainiaoGlobalSolutionServiceResourceQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
