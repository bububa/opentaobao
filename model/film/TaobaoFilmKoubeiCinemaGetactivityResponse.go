package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑-影院营销数据查询 APIResponse
taobao.film.koubei.cinema.getactivity

口碑-影院营销数据查询
*/
type TaobaoFilmKoubeiCinemaGetactivityAPIResponse struct {
    model.CommonResponse
    TaobaoFilmKoubeiCinemaGetactivityResponse
}

type TaobaoFilmKoubeiCinemaGetactivityResponse struct {
    XMLName xml.Name `xml:"film_koubei_cinema_getactivity_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
