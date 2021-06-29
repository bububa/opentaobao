package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票外部直发券 APIResponse
taobao.film.lottery.sendcode

淘票票外部直发券
*/
type TaobaoFilmLotterySendcodeAPIResponse struct {
    model.CommonResponse
    TaobaoFilmLotterySendcodeResponse
}

type TaobaoFilmLotterySendcodeResponse struct {
    XMLName xml.Name `xml:"film_lottery_sendcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
