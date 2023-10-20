package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsmyxgetentinfo 查企业标识信息
// alibaba.alihealth.drug.kyt.smyx.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func Alibabaalihealthdrugkytsmyxgetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsmyxgetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsmyxgetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsmyxgetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
