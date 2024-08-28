package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetentinfo 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
// alibaba.alihealth.drug.kyt.wes.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func AlibabaAlihealthDrugKytWesGetentinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesGetentinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
