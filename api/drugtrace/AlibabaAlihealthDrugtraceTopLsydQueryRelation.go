package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydqueryrelation 单码关联关系查询
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
func Alibabaalihealthdrugtracetoplsydqueryrelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydqueryrelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
