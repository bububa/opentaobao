package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryRelation 单码关联关系查询
// alibaba.alihealth.drugtrace.top.yljg.query.relation
//
// 单码关联关系查询
func AlibabaAlihealthDrugtraceTopYljgQueryRelation(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
