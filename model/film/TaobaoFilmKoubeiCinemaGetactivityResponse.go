package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
口碑-影院营销数据查询 APIResponse
taobao.film.koubei.cinema.getactivity

口碑-影院营销数据查询
*/
type TaobaoFilmKoubeiCinemaGetactivityAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilmKoubeiCinemaGetactivityResponse `json:"film_koubei_cinema_getactivity_response,omitempty"` 
    TaobaoFilmKoubeiCinemaGetactivityResponse
}

/* model for simplify = false
type TaobaoFilmKoubeiCinemaGetactivityResponse struct {

    // result
    
    Result  *struct {
        ResultGeneralModel  *ResultGeneralModel `json:"result_general_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFilmKoubeiCinemaGetactivityResponse struct {

    // result
    Result   *ResultGeneralModel `json:"result,omitempty"`

}
