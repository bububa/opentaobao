package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest
16-供应商预测（OEM-成品）回传接口 API请求
alibaba.tmallgenie.scp.plan.forecast.oem.upload

供应商预测（OEM-成品）回传接口 */
type AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest struct {
	model.Params
	// 请求参数
	_supplierForecastRequest *SupplierForecastRequest
}

// New
