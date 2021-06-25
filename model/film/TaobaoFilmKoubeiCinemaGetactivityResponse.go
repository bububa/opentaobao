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
    Response *TaobaoFilmKoubeiCinemaGetactivityResponse `json:"taobao_film_koubei_cinema_getactivity_response,omitempty"`
}

type TaobaoFilmKoubeiCinemaGetactivityResponse struct {

    // result
    Result   *ResultGeneralModel `json:"result,omitempty"`

}
