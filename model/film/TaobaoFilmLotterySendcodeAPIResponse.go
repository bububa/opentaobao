package film

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmLotterySendcodeAPIResponse 淘票票外部直发券 API返回值
// taobao.film.lottery.sendcode
//
// 淘票票外部直发券
type TaobaoFilmLotterySendcodeAPIResponse struct {
	model.CommonResponse
	TaobaoFilmLotterySendcodeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmLotterySendcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmLotterySendcodeAPIResponseModel).Reset()
}

// TaobaoFilmLotterySendcodeAPIResponseModel is 淘票票外部直发券 成功返回结果
type TaobaoFilmLotterySendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"film_lottery_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmLotterySendcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmLotterySendcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmLotterySendcodeAPIResponse)
	},
}

// GetTaobaoFilmLotterySendcodeAPIResponse 从 sync.Pool 获取 TaobaoFilmLotterySendcodeAPIResponse
func GetTaobaoFilmLotterySendcodeAPIResponse() *TaobaoFilmLotterySendcodeAPIResponse {
	return poolTaobaoFilmLotterySendcodeAPIResponse.Get().(*TaobaoFilmLotterySendcodeAPIResponse)
}

// ReleaseTaobaoFilmLotterySendcodeAPIResponse 将 TaobaoFilmLotterySendcodeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmLotterySendcodeAPIResponse(v *TaobaoFilmLotterySendcodeAPIResponse) {
	v.Reset()
	poolTaobaoFilmLotterySendcodeAPIResponse.Put(v)
}
