package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanforecastoemuploadAPIRequest 16-供应商预测（OEM-成品）回传接口 API请求
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
type AlibabatmallgeniescpplanforecastoemuploadAPIRequest struct {
	model.Params
	// 请求参数
	_supplierForecastRequest *SupplierForecastRequest
}

// NewAlibabatmallgeniescpplanforecastoemuploadRequest 初始化AlibabatmallgeniescpplanforecastoemuploadAPIRequest对象
func NewAlibabatmallgeniescpplanforecastoemuploadRequest() *AlibabatmallgeniescpplanforecastoemuploadAPIRequest {
	return &AlibabatmallgeniescpplanforecastoemuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanforecastoemuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.forecast.oem.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanforecastoemuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanforecastoemuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierForecastRequest is SupplierForecastRequest Setter
// 请求参数
func (r *AlibabatmallgeniescpplanforecastoemuploadAPIRequest) SetSupplierForecastRequest(_supplierForecastRequest *SupplierForecastRequest) error {
	r._supplierForecastRequest = _supplierForecastRequest
	r.Set("supplier_forecast_request", _supplierForecastRequest)
	return nil
}

// GetSupplierForecastRequest SupplierForecastRequest Getter
func (r AlibabatmallgeniescpplanforecastoemuploadAPIRequest) GetSupplierForecastRequest() *SupplierForecastRequest {
	return r._supplierForecastRequest
}
