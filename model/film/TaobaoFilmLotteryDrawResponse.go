package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖发放权益API APIResponse
taobao.film.lottery.draw

对外第三方合作渠道通过抽奖形式发码
*/
type TaobaoFilmLotteryDrawAPIResponse struct {
    model.CommonResponse
    TaobaoFilmLotteryDrawResponse
}

type TaobaoFilmLotteryDrawResponse struct {
    XMLName xml.Name `xml:"film_lottery_draw_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
