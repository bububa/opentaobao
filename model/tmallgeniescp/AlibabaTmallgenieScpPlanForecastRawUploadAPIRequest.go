package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanforecastrawuploadAPIRequest 17-供应商预测（原料-二级物料）回传接口 API请求
// alibaba.tmallgenie.scp.plan.forecast.raw.upload
//
// 供应商预测（原料-二级物料）回传接口
type AlibabatmallgeniescpplanforecastrawuploadAPIRequest struct {
	model.Params
	// 入参对象
	_supplierForecastRawRequest *SupplierForecastRawRequest
}

// NewAlibabatmallgeniescpplanforecastrawuploadRequest 初始化AlibabatmallgeniescpplanforecastrawuploadAPIRequest对象
func NewAlibabatmallgeniescpplanforecastrawuploadRequest() *AlibabatmallgeniescpplanforecastrawuploadAPIRequest {
	return &AlibabatmallgeniescpplanforecastrawuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanforecastrawuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.forecast.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanforecastrawuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanforecastrawuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierForecastRawRequest is SupplierForecastRawRequest Setter
// 入参对象
func (r *AlibabatmallgeniescpplanforecastrawuploadAPIRequest) SetSupplierForecastRawRequest(_supplierForecastRawRequest *SupplierForecastRawRequest) error {
	r._supplierForecastRawRequest = _supplierForecastRawRequest
	r.Set("supplier_forecast_raw_request", _supplierForecastRawRequest)
	return nil
}

// GetSupplierForecastRawRequest SupplierForecastRawRequest Getter
func (r AlibabatmallgeniescpplanforecastrawuploadAPIRequest) GetSupplierForecastRawRequest() *SupplierForecastRawRequest {
	return r._supplierForecastRawRequest
}
