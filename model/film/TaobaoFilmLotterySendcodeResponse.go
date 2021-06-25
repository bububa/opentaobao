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
    Response *TaobaoFilmLotterySendcodeResponse `json:"taobao_film_lottery_sendcode_response,omitempty"`
}

type TaobaoFilmLotterySendcodeResponse struct {

    // result
    Result   *ResultGeneralModel `json:"result,omitempty"`

}
