package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
19-销售预测数量（销管）回传接口 APIRequest
alibaba.tmallgenie.scp.plan.saleforcast.saler.upload

销售预测数量（销管）回传接口
*/
type AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest struct {
    model.Params

    // 入参
    salesForecastRequest   *SalesForecastRequest 

}

func NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest{
    return &AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.saleforcast.saler.upload"
}

func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) SetSalesForecastRequest(salesForecastRequest *SalesForecastRequest) error {
    r.salesForecastRequest = salesForecastRequest
    r.Set("sales_forecast_request", salesForecastRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) GetSalesForecastRequest() *SalesForecastRequest {
    return r.salesForecastRequest
}

