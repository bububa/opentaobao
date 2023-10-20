package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesgetentinfo 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
// alibaba.alihealth.drug.kyt.wes.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func Alibabaalihealthdrugkytwesgetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesgetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesgetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesgetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
