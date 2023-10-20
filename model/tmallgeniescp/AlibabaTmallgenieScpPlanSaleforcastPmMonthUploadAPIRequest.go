package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest 24-销售月预测数量（产管）回传-月度 API请求
// alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload
//
// 销售月预测数量（产管）回传-月度
type AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// NewAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest 初始化AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) Reset() {
	r._salesForecastRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSalesForecastRequest is SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
	r._salesForecastRequest = _salesForecastRequest
	r.Set("sales_forecast_request", _salesForecastRequest)
	return nil
}

// GetSalesForecastRequest SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) GetSalesForecastRequest() *SalesForecastRequest {
	return r._salesForecastRequest
}

var poolAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest
func GetAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest() *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest 将 AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest(v *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest.Put(v)
}
