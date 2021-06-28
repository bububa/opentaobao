package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 APIResponse
taobao.film.data.third.party.refund.order

淘票票第三方退票接口
*/
type TaobaoFilmDataThirdPartyRefundOrderAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilmDataThirdPartyRefundOrderResponse `json:"film_data_third_party_refund_order_response,omitempty"` 
    TaobaoFilmDataThirdPartyRefundOrderResponse
}

/* model for simplify = false
type TaobaoFilmDataThirdPartyRefundOrderResponse struct {

    // result
    
    Result  *struct {
        ResultGeneralModel  *ResultGeneralModel `json:"result_general_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFilmDataThirdPartyRefundOrderResponse struct {

    // result
    Result   *ResultGeneralModel `json:"result,omitempty"`

}
