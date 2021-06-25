package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号v1.0 APIResponse
taobao.wlb.waybill.i.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
type TaobaoWlbWaybillICancelAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWaybillICancelResponse `json:"taobao_wlb_waybill_i_cancel_response,omitempty"`
}

type TaobaoWlbWaybillICancelResponse struct {

    // 调用取消是否成功
    CancelResult   bool `json:"cancel_result,omitempty"`

}
