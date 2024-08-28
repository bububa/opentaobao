package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetentinfo 通过企业名得到唯一标识ref_ent_id及企业ent_id
// alibaba.alihealth.drugtrace.top.yljg.query.getentinfo
//
// 根据企业名称查询ID
func AlibabaAlihealthDrugtraceTopYljgQueryGetentinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
