package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointquerypointflow 分页查询积分流水
// alibaba.alsc.crm.point.querypointflow
//
// 分页查询积分流水
func Alibabaalsccrmpointquerypointflow(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointquerypointflowAPIRequest, session string) (*alsc.AlibabaalsccrmpointquerypointflowAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointquerypointflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
