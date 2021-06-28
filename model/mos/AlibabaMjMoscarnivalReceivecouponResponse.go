package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据手机号码领券 APIResponse
alibaba.mj.moscarnival.receivecoupon

根据手机号码领券
*/
type AlibabaMjMoscarnivalReceivecouponAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjMoscarnivalReceivecouponResponse `json:"alibaba_mj_moscarnival_receivecoupon_response,omitempty"` 
    AlibabaMjMoscarnivalReceivecouponResponse
}

/* model for simplify = false
type AlibabaMjMoscarnivalReceivecouponResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaMjMoscarnivalReceivecouponResultDo  *AlibabaMjMoscarnivalReceivecouponResultDo `json:"alibaba_mj_moscarnival_receivecoupon_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMjMoscarnivalReceivecouponResponse struct {

    // 返回结果
    Result   *AlibabaMjMoscarnivalReceivecouponResultDo `json:"result,omitempty"`

}
