package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydquerygetentinfo 通过企业名得到唯一标识ref_ent_id及企业ent_id
// alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo
//
// 根据企业名称查询ID
func Alibabaalihealthdrugtracetoplsydquerygetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydquerygetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
