package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryBillstatus 上传单据后处理状态查询
// alibaba.alihealth.drugtrace.top.yljg.query.billstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugtraceTopYljgQueryBillstatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
