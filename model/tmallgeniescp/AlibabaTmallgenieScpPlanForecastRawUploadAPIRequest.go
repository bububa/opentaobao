package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest 17-供应商预测（原料-二级物料）回传接口 API请求
// alibaba.tmallgenie.scp.plan.forecast.raw.upload
//
// 供应商预测（原料-二级物料）回传接口
type AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest struct {
	model.Params
	// 入参对象
	_supplierForecastRawRequest *SupplierForecastRawRequest
}

// NewAlibabaTmallgenieScpPlanForecastRawUploadRequest 初始化AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanForecastRawUploadRequest() *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) Reset() {
	r._supplierForecastRawRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.forecast.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierForecastRawRequest is SupplierForecastRawRequest Setter
// 入参对象
func (r *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) SetSupplierForecastRawRequest(_supplierForecastRawRequest *SupplierForecastRawRequest) error {
	r._supplierForecastRawRequest = _supplierForecastRawRequest
	r.Set("supplier_forecast_raw_request", _supplierForecastRawRequest)
	return nil
}

// GetSupplierForecastRawRequest SupplierForecastRawRequest Getter
func (r AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) GetSupplierForecastRawRequest() *SupplierForecastRawRequest {
	return r._supplierForecastRawRequest
}

var poolAlibabaTmallgenieScpPlanForecastRawUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanForecastRawUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanForecastRawUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest
func GetAlibabaTmallgenieScpPlanForecastRawUploadAPIRequest() *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanForecastRawUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanForecastRawUploadAPIRequest 将 AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanForecastRawUploadAPIRequest(v *AlibabaTmallgenieScpPlanForecastRawUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanForecastRawUploadAPIRequest.Put(v)
}
