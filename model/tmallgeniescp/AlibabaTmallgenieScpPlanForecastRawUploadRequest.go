package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
17-供应商预测（原料-二级物料）回传接口 API请求
alibaba.tmallgenie.scp.plan.forecast.raw.upload

供应商预测（原料-二级物料）回传接口
*/
type AlibabaTmallgenieScpPlanForecastRawUploadRequest struct {
    model.Params
    // 入参对象
    supplierForecastRawRequest   *SupplierForecastRawRequest
}

// 初始化AlibabaTmallgenieScpPlanForecastRawUploadRequest对象
func NewAlibabaTmallgenieScpPlanForecastRawUploadRequest() *AlibabaTmallgenieScpPlanForecastRawUploadRequest{
    return &AlibabaTmallgenieScpPlanForecastRawUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanForecastRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.forecast.raw.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanForecastRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierForecastRawRequest Setter
// 入参对象
func (r *AlibabaTmallgenieScpPlanForecastRawUploadRequest) SetSupplierForecastRawRequest(supplierForecastRawRequest *SupplierForecastRawRequest) error {
    r.supplierForecastRawRequest = supplierForecastRawRequest
    r.Set("supplier_forecast_raw_request", supplierForecastRawRequest)
    return nil
}

// SupplierForecastRawRequest Getter
func (r AlibabaTmallgenieScpPlanForecastRawUploadRequest) GetSupplierForecastRawRequest() *SupplierForecastRawRequest {
    return r.supplierForecastRawRequest
}
