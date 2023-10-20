package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalsolutionserviceresourcequery 查询解决方案服务资源列表
// cainiao.global.solution.service.resource.query
//
// 返回直接解决方案的指定物流服务的可用资源列表
func Cainiaoglobalsolutionserviceresourcequery(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalsolutionserviceresourcequeryAPIRequest, session string) (*cainiaohandover.CainiaoglobalsolutionserviceresourcequeryAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalsolutionserviceresourcequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
