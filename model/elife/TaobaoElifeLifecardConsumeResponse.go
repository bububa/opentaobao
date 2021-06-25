package elife

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券核销 APIResponse
taobao.elife.lifecard.consume

用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
*/
type TaobaoElifeLifecardConsumeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoElifeLifecardConsumeResponse `json:"taobao_elife_lifecard_consume_response,omitempty"`
}

type TaobaoElifeLifecardConsumeResponse struct {

    // 本金
    Amount   int64 `json:"amount,omitempty"`

    // 膨胀金
    InflateAmount   int64 `json:"inflate_amount,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 结果描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 是否成功
    Successed   bool `json:"successed,omitempty"`

}
