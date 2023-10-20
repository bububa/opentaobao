package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgquerylistparts 往来单位查询
// alibaba.alihealth.drugtrace.top.yljg.query.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugtracetopyljgquerylistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgquerylistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
