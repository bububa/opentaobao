package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytyygetentinfo 得到企业信息
// alibaba.alihealth.drug.kyt.yy.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func Alibabaalihealthdrugkytyygetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytyygetentinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytyygetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytyygetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
