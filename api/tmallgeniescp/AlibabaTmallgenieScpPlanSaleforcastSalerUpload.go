package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSaleforcastSalerUpload 19-销售预测数量（销管）回传接口
// alibaba.tmallgenie.scp.plan.saleforcast.saler.upload
//
// 销售预测数量（销管）回传接口
func AlibabaTmallgenieScpPlanSaleforcastSalerUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
