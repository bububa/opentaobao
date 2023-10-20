package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixopengatewayprojectstatusquery 分销状态查询接口queryProjectStatusByProjectId
// alibaba.damai.maitix.opengateway.project.status.query
//
// queryProjectStatusByProjectId
func Alibabadamaimaitixopengatewayprojectstatusquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixopengatewayprojectstatusqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixopengatewayprojectstatusqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
