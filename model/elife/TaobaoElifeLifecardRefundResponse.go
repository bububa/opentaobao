package elife

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券冲正退还 APIResponse
taobao.elife.lifecard.refund

淘宝生活汇消费卡虚拟卡，线下冲正退货接口
*/
type TaobaoElifeLifecardRefundAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoElifeLifecardRefundResponse `json:"elife_lifecard_refund_response,omitempty"` 
    TaobaoElifeLifecardRefundResponse
}

/* model for simplify = false
type TaobaoElifeLifecardRefundResponse struct {

    // 返回码，成功为空
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 本金
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 膨胀金
    
    InflateAmount   int64 `json:"inflate_amount,omitempty"`
    

    // 返回信息
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 成功失败标志
    
    Successed   bool `json:"successed,omitempty"`
    

}
*/

type TaobaoElifeLifecardRefundResponse struct {

    // 返回码，成功为空
    ResultCode   string `json:"result_code,omitempty"`

    // 本金
    Amount   int64 `json:"amount,omitempty"`

    // 膨胀金
    InflateAmount   int64 `json:"inflate_amount,omitempty"`

    // 返回信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // 成功失败标志
    Successed   bool `json:"successed,omitempty"`

}
