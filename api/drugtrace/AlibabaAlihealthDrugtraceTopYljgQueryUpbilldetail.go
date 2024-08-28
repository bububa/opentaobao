package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetail 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】
// alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
func AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
