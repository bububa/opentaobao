package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowDrugSupervisionQuery 监管平台药品查询
// alibaba.alihealth.outflow.drug.supervision.query
//
// 获取监管平台药品数据
func AlibabaAlihealthOutflowDrugSupervisionQuery(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
