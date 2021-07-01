package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest
17-供应商预测（原料-二级物料）回传接口 API请求
alibaba.tmallgenie.scp.plan.forecast.raw.upload

供应商预测（原料-二级物料）回传接口 */
type AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest struct {
	model.Params
	// 入参对象
	_supplierForecastRawRequest *SupplierForecastRawRequest
}

// New
