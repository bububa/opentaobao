package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
18-销售预测数量（产管）回传接口 APIRequest
alibaba.tmallgenie.scp.plan.saleforcast.pm.upload

销售预测数量（产管）回传接口
*/
type AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest struct {
    model.Params

    // 入参
    salesForecastRequest   *SalesForecastRequest 

}

func NewAlibabaTmallgenieScpPlanSaleforcastPmUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest{
    return &AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.saleforcast.pm.upload"
}

func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) SetSalesForecastRequest(salesForecastRequest *SalesForecastRequest) error {
    r.salesForecastRequest = salesForecastRequest
    r.Set("sales_forecast_request", salesForecastRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest) GetSalesForecastRequest() *SalesForecastRequest {
    return r.salesForecastRequest
}

