package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
24-销售月预测数量（产管）回传-月度 API请求
alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload

销售月预测数量（产管）回传-月度
*/
type AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest struct {
    model.Params
    // 入参
    salesForecastRequest   *SalesForecastRequest
}

// 初始化AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest{
    return &AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) SetSalesForecastRequest(salesForecastRequest *SalesForecastRequest) error {
    r.salesForecastRequest = salesForecastRequest
    r.Set("sales_forecast_request", salesForecastRequest)
    return nil
}

// SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) GetSalesForecastRequest() *SalesForecastRequest {
    return r.salesForecastRequest
}
