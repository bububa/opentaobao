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
	RequestId     string         `json:"request_id,omitempty" xml:"film_lottery_sendcode_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultGeneralModel `json:"result,omitempty" xml:"