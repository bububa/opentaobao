package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘票票外部直发券 APIResponse
taobao.film.lottery.sendcode

淘票票外部直发券
*/
type TaobaoFilmLotterySendcodeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilmLotterySendcodeResponse `json:"film_lottery_sendcode_response,omitempty"` 
    TaobaoFilmLotterySendcodeResponse
}

/* model for simplify = false
type TaobaoFilmLotterySendcodeResponse struct {

    // result
    
    Result  *struct {
        ResultGeneralModel  *ResultGeneralModel `json:"result_general_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFilmLotterySendcodeResponse struct {

    // result
    Result   *ResultGeneralModel `json:"result,omitempty"`

}
