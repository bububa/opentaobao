package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgListupoutDetail 上游出库单单据明细查询
// alibaba.alihealth.drugtrace.top.yljg.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugtraceTopYljgListupoutDetail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
