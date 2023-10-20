package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrgetentinfo 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id)
// alibaba.alihealth.drug.kyt.dr.getentinfo
//
// 根据企业名称查询ID
func Alibabaalihealthdrugkytdrgetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrgetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrgetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrgetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
