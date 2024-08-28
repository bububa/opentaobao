package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeAdvanceBillFlowDirection 单据流向查询
// alibaba.alihealth.drug.code.advance.bill.flow.direction
//
// 单据流向查询
func AlibabaAlihealthDrugCodeAdvanceBillFlowDirection(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
