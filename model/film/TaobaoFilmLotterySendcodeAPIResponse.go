package film

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmlotterysendcodeAPIResponse 淘票票外部直发券 API返回值
// taobao.film.lottery.sendcode
//
// 淘票票外部直发券
type TaobaofilmlotterysendcodeAPIResponse struct {
	model.CommonResponse
	TaobaofilmlotterysendcodeAPIResponseModel
}

// TaobaofilmlotterysendcodeAPIResponseModel is 淘票票外部直发券 成功返回结果
type TaobaofilmlotterysendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"film_lottery_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`
}
