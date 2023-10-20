package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanForecastRawUpload 17-供应商预测（原料-二级物料）回传接口
// alibaba.tmallgenie.scp.plan.forecast.raw.upload
//
// 供应商预测（原料-二级物料）回传接口
func AlibabaTmallgenieScpPlanForecastRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanForecastRawUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
