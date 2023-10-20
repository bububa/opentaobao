package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetentinfo 通过企业名得到唯一标识(ref_ent_id)及企业ID(ent_id)
// alibaba.alihealth.drug.kyt.specia.vaccin.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
func AlibabaAlihealthDrugKytSpeciaVaccinGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
