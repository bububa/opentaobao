package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytVaGetentinfo
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id)
alibaba.alihealth.drug.kyt.va.getentinfo

根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id) */
func AlibabaAlihealthDrugKytVaGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytVaGetentinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytVaGetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytVaGetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
