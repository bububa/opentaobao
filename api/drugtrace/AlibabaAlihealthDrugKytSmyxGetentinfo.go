package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSmyxGetentinfo 查企业标识信息
// alibaba.alihealth.drug.kyt.smyx.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
func AlibabaAlihealthDrugKytSmyxGetentinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
