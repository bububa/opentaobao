package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetentinfo 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func Alibabaalihealthdrugkytgetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
