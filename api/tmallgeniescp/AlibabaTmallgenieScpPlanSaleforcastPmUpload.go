package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSaleforcastPmUpload 18-销售预测数量（产管）回传接口
// alibaba.tmallgenie.scp.plan.saleforcast.pm.upload
//
// 销售预测数量（产管）回传接口
func AlibabaTmallgenieScpPlanSaleforcastPmUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
