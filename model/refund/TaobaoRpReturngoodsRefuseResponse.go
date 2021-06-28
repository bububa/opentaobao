package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退货 APIResponse
taobao.rp.returngoods.refuse

卖家拒绝退货，目前仅支持天猫退货。
*/
type TaobaoRpReturngoodsRefuseAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRpReturngoodsRefuseResponse `json:"rp_returngoods_refuse_response,omitempty"` 
    TaobaoRpReturngoodsRefuseResponse
}

/* model for simplify = false
type TaobaoRpReturngoodsRefuseResponse struct {

    // asdf
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoRpReturngoodsRefuseResponse struct {

    // asdf
    Result   bool `json:"result,omitempty"`

}
