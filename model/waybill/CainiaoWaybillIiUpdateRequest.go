package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印更新接口 APIRequest
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。
*/
type CainiaoWaybillIiUpdateRequest struct {
    model.Params

    // 更新请求信息
    paramWaybillCloudPrintUpdateRequest   *WaybillCloudPrintUpdateRequest 

}

func NewCainiaoWaybillIiUpdateRequest() *CainiaoWaybillIiUpdateRequest{
    return &CainiaoWaybillIiUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiUpdateRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.update"
}

func (r CainiaoWaybillIiUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiUpdateRequest) SetParamWaybillCloudPrintUpdateRequest(paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest) error {
    r.paramWaybillCloudPrintUpdateRequest = paramWaybillCloudPrintUpdateRequest
    r.Set("param_waybill_cloud_print_update_request", paramWaybillCloudPrintUpdateRequest)
    return nil
}

func (r CainiaoWaybillIiUpdateRequest) GetParamWaybillCloudPrintUpdateRequest() *WaybillCloudPrintUpdateRequest {
    return r.paramWaybillCloudPrintUpdateRequest
}

