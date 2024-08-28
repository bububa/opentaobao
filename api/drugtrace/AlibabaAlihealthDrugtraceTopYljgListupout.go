package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgListupout 医疗机构查询本企业上游企业出库单据信息
// alibaba.alihealth.drugtrace.top.yljg.listupout
//
// 查询货主/本企业上游企业出库单据信息
func AlibabaAlihealthDrugtraceTopYljgListupout(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
