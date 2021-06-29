package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
19-销售预测数量（销管）回传接口 API请求
alibaba.tmallgenie.scp.plan.saleforcast.saler.upload

销售预测数量（销管）回传接口
*/
type AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest struct {
    model.Params
    // 入参
    _salesForecastRequest   *SalesForecastRequest
}

// 初始化AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest{
    return &AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.saleforcast.saler.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
    r._salesForecastRequest = _salesForecastRequest
    r.Set("sales_forecast_request", _salesForecastRequest)
    return nil
}

// SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest) GetSalesForecastRequest() *SalesForecastRequest {
    return r._salesForecastRequest
}
