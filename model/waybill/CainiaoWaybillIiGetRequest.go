package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印接口 API请求
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法
*/
type CainiaoWaybillIiGetRequest struct {
    model.Params
    // 入参信息
    _paramWaybillCloudPrintApplyNewRequest   *WaybillCloudPrintApplyNewRequest
}

// 初始化CainiaoWaybillIiGetRequest对象
func NewCainiaoWaybillIiGetRequest() *CainiaoWaybillIiGetRequest{
    return &CainiaoWaybillIiGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamWaybillCloudPrintApplyNewRequest Setter
// 入参信息
func (r *CainiaoWaybillIiGetRequest) SetParamWaybillCloudPrintApplyNewRequest(_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest) error {
    r._paramWaybillCloudPrintApplyNewRequest = _paramWaybillCloudPrintApplyNewRequest
    r.Set("param_waybill_cloud_print_apply_new_request", _paramWaybillCloudPrintApplyNewRequest)
    return nil
}

// ParamWaybillCloudPrintApplyNewRequest Getter
func (r CainiaoWaybillIiGetRequest) GetParamWaybillCloudPrintApplyNewRequest() *WaybillCloudPrintApplyNewRequest {
    return r._paramWaybillCloudPrintApplyNewRequest
}
