package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrGetentinfo 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id)
// alibaba.alihealth.drug.kyt.dr.getentinfo
//
// 根据企业名称查询ID
func AlibabaAlihealthDrugKytDrGetentinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrGetentinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
