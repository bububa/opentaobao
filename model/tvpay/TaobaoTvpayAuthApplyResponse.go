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
    Response *TaobaoTvpayAuthApplyResponse `json:"taobao_tvpay_auth_apply_response,omitempty"`
}

type TaobaoTvpayAuthApplyResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
