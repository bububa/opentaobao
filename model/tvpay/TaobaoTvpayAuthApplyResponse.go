package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付申请设备授权 APIResponse
taobao.tvpay.auth.apply

为用户在指定设备上申请支付授权
*/
type TaobaoTvpayAuthApplyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTvpayAuthApplyResponse `json:"tvpay_auth_apply_response,omitempty"` 
    TaobaoTvpayAuthApplyResponse
}

/* model for simplify = false
type TaobaoTvpayAuthApplyResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayAuthApplyResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
