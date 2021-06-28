package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号 APIResponse
cainiao.waybill.ii.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
type CainiaoWaybillIiCancelAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillIiCancelResponse `json:"cainiao_waybill_ii_cancel_response,omitempty"` 
    CainiaoWaybillIiCancelResponse
}

/* model for simplify = false
type CainiaoWaybillIiCancelResponse struct {

    // 调用取消是否成功
    
    CancelResult   bool `json:"cancel_result,omitempty"`
    

}
*/

type CainiaoWaybillIiCancelResponse struct {

    // 调用取消是否成功
    CancelResult   bool `json:"cancel_result,omitempty"`

}
