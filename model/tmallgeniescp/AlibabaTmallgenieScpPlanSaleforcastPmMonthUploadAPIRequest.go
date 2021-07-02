package tmallgeniescp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
