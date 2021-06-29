package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印更新接口 API请求
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。
*/
type CainiaoWaybillIiUpdateRequest struct {
    model.Params
    // 更新请求信息
    paramWaybillCloudPrintUpdateRequest   *WaybillCloudPrintUpdateRequest
}

// 初始化CainiaoWaybillIiUpdateRequest对象
func NewCainiaoWaybillIiUpdateRequest() *CainiaoWaybillIiUpdateRequest{
    return &CainiaoWaybillIiUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiUpdateRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.update"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamWaybillCloudPrintUpdateRequest Setter
// 更新请求信息
func (r *CainiaoWaybillIiUpdateRequest) SetParamWaybillCloudPrintUpdateRequest(paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest) error {
    r.paramWaybillCloudPrintUpdateRequest = paramWaybillCloudPrintUpdateRequest
    r.Set("param_waybill_cloud_print_update_request", paramWaybillCloudPrintUpdateRequest)
    return nil
}

// ParamWaybillCloudPrintUpdateRequest Getter
func (r CainiaoWaybillIiUpdateRequest) GetParamWaybillCloudPrintUpdateRequest() *WaybillCloudPrintUpdateRequest {
    return r.paramWaybillCloudPrintUpdateRequest
}
