package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetentinfo 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
// alibaba.alihealth.drug.kyt.wes.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func AlibabaAlihealthDrugKytWesGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetentinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesGetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesGetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
