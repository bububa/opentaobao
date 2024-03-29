package film

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse CGV影城卡排期数据传输 API返回值
// taobao.film.tfbackyard.cardschedule.update
//
// cgv影城卡排期价格数据传输API
type TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfbackyardCardscheduleUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmTfbackyardCardscheduleUpdateAPIResponseModel).Reset()
}

// TaobaoFilmTfbackyardCardscheduleUpdateAPIResponseModel is CGV影城卡排期数据传输 成功返回结果
type TaobaoFilmTfbackyardCardscheduleUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfbackyard_cardschedule_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmTfbackyardCardscheduleUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse)
	},
}

// GetTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse 从 sync.Pool 获取 TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse
func GetTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse() *TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse {
	return poolTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse.Get().(*TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse)
}

// ReleaseTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse 将 TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse(v *TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFilmTfbackyardCardscheduleUpdateAPIResponse.Put(v)
}
