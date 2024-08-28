package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthAsyncprescribePrescriptionDetail 异步开方处方详情
// alibaba.alihealth.asyncprescribe.prescription.detail
//
// 异步开方处方查询
func AlibabaAlihealthAsyncprescribePrescriptionDetail(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest, resp *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
