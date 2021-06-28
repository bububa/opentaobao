package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退款信息审核 APIResponse
taobao.rdc.aligenius.refunds.check

根据退款信息，对退款单进行审核
*/
type TaobaoRdcAligeniusRefundsCheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdcAligeniusRefundsCheckResponse `json:"rdc_aligenius_refunds_check_response,omitempty"` 
    TaobaoRdcAligeniusRefundsCheckResponse
}

/* model for simplify = false
type TaobaoRdcAligeniusRefundsCheckResponse struct {

    // result
    
    Result  *struct {
        TaobaoRdcAligeniusRefundsCheckResult  *TaobaoRdcAligeniusRefundsCheckResult `json:"taobao_rdc_aligenius_refunds_check_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdcAligeniusRefundsCheckResponse struct {

    // result
    Result   *TaobaoRdcAligeniusRefundsCheckResult `json:"result,omitempty"`

}
