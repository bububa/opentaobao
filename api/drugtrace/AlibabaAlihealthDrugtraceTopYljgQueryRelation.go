package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgqueryrelation 单码关联关系查询
// alibaba.alihealth.drugtrace.top.yljg.query.relation
//
// 单码关联关系查询
func Alibabaalihealthdrugtracetopyljgqueryrelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgqueryrelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgqueryrelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgqueryrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
