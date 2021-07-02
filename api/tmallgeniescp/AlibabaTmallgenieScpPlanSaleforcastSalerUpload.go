package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanSaleforcastSalerUpload 19-销售预测数量（销管）回传接口
// alibaba.tmallgenie.scp.plan.saleforcast.saler.upload
//
// 销售预测数量（销管）回传接口
func AlibabaTmallgenieScpPlanSaleforcastSalerUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
