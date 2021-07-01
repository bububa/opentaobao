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

// NewAlibabaTmallgenieScpPlanForecastOemUploadRequest 初始化AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanForecastOemUploadRequest() *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.forecast.oem.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SupplierForecastRequest Setter
// 请求参数
func (r *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) SetSupplierForecastRequest(_supplierForecastRequest *SupplierForecastRequest) error {
	r._supplierForecastRequest = _supplierForecastRequest
	r.Set("supplier_forecast_request", _supplierForecastRequest)
	return nil
}

// Get SupplierForecastRequest Getter
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetSupplierForecastRequest() *SupplierForecastRequest {
	return r._supplierForecastRequest
}
