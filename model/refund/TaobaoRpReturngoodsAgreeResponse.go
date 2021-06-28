package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意退货 APIResponse
taobao.rp.returngoods.agree

卖家同意退货，支持淘宝和天猫的订单。
*/
type TaobaoRpReturngoodsAgreeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRpReturngoodsAgreeResponse `json:"rp_returngoods_agree_response,omitempty"` 
    TaobaoRpReturngoodsAgreeResponse
}

/* model for simplify = false
type TaobaoRpReturngoodsAgreeResponse struct {

    // 操作成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoRpReturngoodsAgreeResponse struct {

    // 操作成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
