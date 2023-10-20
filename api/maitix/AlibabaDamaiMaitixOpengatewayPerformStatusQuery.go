package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixopengatewayperformstatusquery 分销状态查询接口queryPerformStatusByPerformId
// alibaba.damai.maitix.opengateway.perform.status.query
//
// queryPerformStatusByPerformId
func Alibabadamaimaitixopengatewayperformstatusquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixopengatewayperformstatusqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixopengatewayperformstatusqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
