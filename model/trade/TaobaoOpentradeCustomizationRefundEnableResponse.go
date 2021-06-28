package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定制订单设置允许仅退款 APIResponse
taobao.opentrade.customization.refund.enable

定制订单设置允许仅退款
*/
type TaobaoOpentradeCustomizationRefundEnableAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpentradeCustomizationRefundEnableResponse `json:"opentrade_customization_refund_enable_response,omitempty"` 
    TaobaoOpentradeCustomizationRefundEnableResponse
}

/* model for simplify = false
type TaobaoOpentradeCustomizationRefundEnableResponse struct {

    // 是否设置成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoOpentradeCustomizationRefundEnableResponse struct {

    // 是否设置成功
    Result   bool `json:"result,omitempty"`

}
