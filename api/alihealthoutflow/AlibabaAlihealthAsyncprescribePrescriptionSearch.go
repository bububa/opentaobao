package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthAsyncprescribePrescriptionSearch 异步开方处方查询
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
func AlibabaAlihealthAsyncprescribePrescriptionSearch(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest, resp *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
