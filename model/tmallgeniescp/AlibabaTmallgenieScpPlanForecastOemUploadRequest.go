package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
16-供应商预测（OEM-成品）回传接口 API请求
alibaba.tmallgenie.scp.plan.forecast.oem.upload

供应商预测（OEM-成品）回传接口
*/
type AlibabaTmallgenieScpPlanForecastOemUploadRequest struct {
    model.Params
    // 请求参数
    _supplierForecastRequest   *SupplierForecastRequest
}

// 初始化AlibabaTmallgenieScpPlanForecastOemUploadRequest对象
func NewAlibabaTmallgenieScpPlanForecastOemUploadRequest() *AlibabaTmallgenieScpPlanForecastOemUploadRequest{
    return &AlibabaTmallgenieScpPlanForecastOemUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanForecastOemUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.forecast.oem.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanForecastOemUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierForecastRequest Setter
// 请求参数
func (r *AlibabaTmallgenieScpPlanForecastOemUploadRequest) SetSupplierForecastRequest(_supplierForecastRequest *SupplierForecastRequest) error {
    r._supplierForecastRequest = _supplierForecastRequest
    r.Set("supplier_forecast_request", _supplierForecastRequest)
    return nil
}

// SupplierForecastRequest Getter
func (r AlibabaTmallgenieScpPlanForecastOemUploadRequest) GetSupplierForecastRequest() *SupplierForecastRequest {
    return r._supplierForecastRequest
}
