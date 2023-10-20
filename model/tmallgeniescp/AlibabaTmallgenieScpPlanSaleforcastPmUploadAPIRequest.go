package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest 18-销售预测数量（产管）回传接口 API请求
// alibaba.tmallgenie.scp.plan.saleforcast.pm.upload
//
// 销售预测数量（产管）回传接口
type AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// NewAlibabaTmallgenieScpPlanSaleforcastPmUploadRequest 初始化AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastPmUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) Reset() {
	r._salesForecastRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.pm.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSalesForecastRequest is SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
	r._salesForecastRequest = _salesForecastRequest
	r.Set("sales_forecast_request", _salesForecastRequest)
	return nil
}

// GetSalesForecastRequest SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) GetSalesForecastRequest() *SalesForecastRequest {
	return r._salesForecastRequest
}

var poolAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanSaleforcastPmUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanSaleforcastPmUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest
func GetAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest() *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest 将 AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest(v *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest.Put(v)
}
