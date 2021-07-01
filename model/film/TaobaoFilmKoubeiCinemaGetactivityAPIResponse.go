package film

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmKoubeiCinemaGetactivityAPIResponse
口碑-影院营销数据查询 API返回值
taobao.film.koubei.cinema.getactivity

口碑-影院营销数据查询 */
type TaobaoFilmKoubeiCinemaGetactivityAPIResponse struct {
	model.CommonResponse
	TaobaoFilmKoubeiCinemaGetactivityAPIResponseModel
}

// TaobaoFilmKoubeiCinemaGetactivityAPIResponseModel is 口碑-影院营销数据查询 成功返回结果
type TaobaoFilmKoubeiCinemaGetactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"film_koubei_cinema_getactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`
}
