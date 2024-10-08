package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbillcode 通过一个码，查询这个码对应的上游企业出库单的单据号
// alibaba.alihealth.drugtrace.top.yljg.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
func AlibabaAlihealthDrugtraceTopYljgQueryUpbillcode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
