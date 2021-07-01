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

// NewAlibabaTmallgenieScpPlanForecastRawUploadRequest 初始化AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanForecastRawUploadRequest() *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.forecast.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SupplierForecastRawRequest Setter
// 入参对象
func (r *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) SetSupplierForecastRawRequest(_supplierForecastRawRequest *SupplierForecastRawRequest) error {
	r._supplierForecastRawRequest = _supplierForecastRawRequest
	r.Set("supplier_forecast_raw_request", _supplierForecastRawRequest)
	return nil
}

// Get SupplierForecastRawRequest Getter
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetSupplierForecastRawRequest() *SupplierForecastRawRequest {
	return r._supplierForecastRawRequest
}
