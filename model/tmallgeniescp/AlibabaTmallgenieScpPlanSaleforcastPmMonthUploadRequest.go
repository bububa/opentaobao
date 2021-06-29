package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
24-销售月预测数量（产管）回传-月度 APIRequest
alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload

销售月预测数量（产管）回传-月度
*/
type AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest struct {
    model.Params

    // 入参
    salesForecastRequest   *SalesForecastRequest 

}

func NewAlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest{
    return &AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload"
}

func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) SetSalesForecastRequest(salesForecastRequest *SalesForecastRequest) error {
    r.salesForecastRequest = salesForecastRequest
    r.Set("sales_forecast_request", salesForecastRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest) GetSalesForecastRequest() *SalesForecastRequest {
    return r.salesForecastRequest
}

