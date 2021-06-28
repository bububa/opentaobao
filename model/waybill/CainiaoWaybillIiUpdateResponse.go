package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印更新接口 APIResponse
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。
*/
type CainiaoWaybillIiUpdateAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillIiUpdateResponse `json:"cainiao_waybill_ii_update_response,omitempty"` 
    CainiaoWaybillIiUpdateResponse
}

/* model for simplify = false
type CainiaoWaybillIiUpdateResponse struct {

    // 面单号
    
    WaybillCode   string `json:"waybill_code,omitempty"`
    

    // 模板内容
    
    PrintData   string `json:"print_data,omitempty"`
    

}
*/

type CainiaoWaybillIiUpdateResponse struct {

    // 面单号
    WaybillCode   string `json:"waybill_code,omitempty"`

    // 模板内容
    PrintData   string `json:"print_data,omitempty"`

}
