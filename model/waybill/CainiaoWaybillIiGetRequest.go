package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印接口 APIRequest
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法
*/
type CainiaoWaybillIiGetRequest struct {
    model.Params

    // 入参信息
    paramWaybillCloudPrintApplyNewRequest   *WaybillCloudPrintApplyNewRequest 

}

func NewCainiaoWaybillIiGetRequest() *CainiaoWaybillIiGetRequest{
    return &CainiaoWaybillIiGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.get"
}

func (r CainiaoWaybillIiGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiGetRequest) SetParamWaybillCloudPrintApplyNewRequest(paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
    r.paramWaybillCloudPrintApplyNewRequest = paramWaybillCloudPrintApplyNewRequest
    r.Set("param_waybill_cloud_print_apply_new_request", paramWaybillCloudPrintApplyNewRequest)
    return nil
}

func (r CainiaoWaybillIiGetRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
    return r.paramWaybillCloudPrintApplyNewRequest
}

