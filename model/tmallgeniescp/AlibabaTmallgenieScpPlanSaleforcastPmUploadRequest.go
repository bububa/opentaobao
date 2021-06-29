package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
18-销售预测数量（产管）回传接口 API请求
alibaba.tmallgenie.scp.plan.saleforcast.pm.upload

销售预测数量（产管）回传接口
*/
type AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest struct {
    model.Params
    // 入参
    salesForecastRequest   *SalesForecastRequest
}

// 初始化AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastPmUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest{
    return &AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.saleforcast.pm.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) SetSalesForecastRequest(salesForecastRequest *SalesForecastRequest) error {
    r.salesForecastRequest = salesForecastRequest
    r.Set("sales_forecast_request", salesForecastRequest)
    return nil
}

// SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) GetSalesForecastRequest() *SalesForecastRequest {
    return r.salesForecastRequest
}
