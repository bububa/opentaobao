package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
16-供应商预测（OEM-成品）回传接口 APIRequest
alibaba.tmallgenie.scp.plan.forecast.oem.upload

供应商预测（OEM-成品）回传接口
*/
type AlibabaTmallgenieScpPlanForecastOemUploadRequest struct {
    model.Params

    // 请求参数
    supplierForecastRequest   *SupplierForecastRequest 

}

func NewAlibabaTmallgenieScpPlanForecastOemUploadRequest() *AlibabaTmallgenieScpPlanForecastOemUploadRequest{
    return &AlibabaTmallgenieScpPlanForecastOemUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanForecastOemUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.forecast.oem.upload"
}

func (r AlibabaTmallgenieScpPlanForecastOemUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanForecastOemUploadRequest) SetSupplierForecastRequest(supplierForecastRequest *SupplierForecastRequest) error {
    r.supplierForecastRequest = supplierForecastRequest
    r.Set("supplier_forecast_request", supplierForecastRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanForecastOemUploadRequest) GetSupplierForecastRequest() *SupplierForecastRequest {
    return r.supplierForecastRequest
}

