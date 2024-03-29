package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest 16-供应商预测（OEM-成品）回传接口 API请求
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
type AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest struct {
	model.Params
	// 请求参数
	_supplierForecastRequest *SupplierForecastRequest
}

// NewAlibabaTmallgenieScpPlanForecastOemUploadRequest 初始化AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanForecastOemUploadRequest() *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) Reset() {
	r._supplierForecastRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.forecast.oem.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierForecastRequest is SupplierForecastRequest Setter
// 请求参数
func (r *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) SetSupplierForecastRequest(_supplierForecastRequest *SupplierForecastRequest) error {
	r._supplierForecastRequest = _supplierForecastRequest
	r.Set("supplier_forecast_request", _supplierForecastRequest)
	return nil
}

// GetSupplierForecastRequest SupplierForecastRequest Getter
func (r AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) GetSupplierForecastRequest() *SupplierForecastRequest {
	return r._supplierForecastRequest
}

var poolAlibabaTmallgenieScpPlanForecastOemUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanForecastOemUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanForecastOemUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest
func GetAlibabaTmallgenieScpPlanForecastOemUploadAPIRequest() *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanForecastOemUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanForecastOemUploadAPIRequest 将 AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanForecastOemUploadAPIRequest(v *AlibabaTmallgenieScpPlanForecastOemUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanForecastOemUploadAPIRequest.Put(v)
}
