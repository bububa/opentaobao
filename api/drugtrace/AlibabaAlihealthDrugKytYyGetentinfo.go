package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyGetentinfo 得到企业信息
// alibaba.alihealth.drug.kyt.yy.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func AlibabaAlihealthDrugKytYyGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyGetentinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
