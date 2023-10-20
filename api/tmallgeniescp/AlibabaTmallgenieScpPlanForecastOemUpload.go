package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanForecastOemUpload 16-供应商预测（OEM-成品）回传接口
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
func AlibabaTmallgenieScpPlanForecastOemUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
