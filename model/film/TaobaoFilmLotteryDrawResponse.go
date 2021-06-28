package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖发放权益API APIResponse
taobao.film.lottery.draw

对外第三方合作渠道通过抽奖形式发码
*/
type TaobaoFilmLotteryDrawAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilmLotteryDrawResponse `json:"film_lottery_draw_response,omitempty"` 
    TaobaoFilmLotteryDrawResponse
}

/* model for simplify = false
type TaobaoFilmLotteryDrawResponse struct {

    // 返回值
    
    Result  *struct {
        ResultGeneralModel  *ResultGeneralModel `json:"result_general_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFilmLotteryDrawResponse struct {

    // 返回值
    Result   *ResultGeneralModel `json:"result,omitempty"`

}
