package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanForecastRawUpload 17-供应商预测（原料-二级物料）回传接口
// alibaba.tmallgenie.scp.plan.forecast.raw.upload
//
// 供应商预测（原料-二级物料）回传接口
func AlibabaTmallgenieScpPlanForecastRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanForecastRawUploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanForecastRawUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
