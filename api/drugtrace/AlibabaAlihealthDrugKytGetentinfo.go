package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetentinfo 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func AlibabaAlihealthDrugKytGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetentinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
