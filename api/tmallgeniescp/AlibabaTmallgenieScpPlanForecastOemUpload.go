package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanForecastOemUpload 16-供应商预测（OEM-成品）回传接口
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
func AlibabaTmallgenieScpPlanForecastOemUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
