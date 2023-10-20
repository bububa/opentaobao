package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydquerylistparts 往来单位查询
// alibaba.alihealth.drugtrace.top.lsyd.query.listparts
//
// 查询往来单位列表
func Alibabaalihealthdrugtracetoplsydquerylistparts(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydquerylistpartsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydquerylistpartsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydquerylistpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
