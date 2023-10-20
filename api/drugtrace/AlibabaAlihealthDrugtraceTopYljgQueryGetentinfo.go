package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgquerygetentinfo 通过企业名得到唯一标识ref_ent_id及企业ent_id
// alibaba.alihealth.drugtrace.top.yljg.query.getentinfo
//
// 根据企业名称查询ID
func Alibabaalihealthdrugtracetopyljgquerygetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgquerygetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
