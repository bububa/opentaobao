package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmallgenieHotelwelcomeAPIResponse 酒店欢迎词推送 API返回值
// taobao.tmallgenie.hotelwelcome
//
// 推送欢迎词，让天猫精灵播放
type TaobaoTmallgenieHotelwelcomeAPIResponse struct {
	model.CommonResponse
	TaobaoTmallgenieHotelwelcomeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmallgenieHotelwelcomeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmallgenieHotelwelcomeAPIResponseModel).Reset()
}

// TaobaoTmallgenieHotelwelcomeAPIResponseModel is 酒店欢迎词推送 成功返回结果
type TaobaoTmallgenieHotelwelcomeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmallgenie_hotelwelcome_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmallgenieHotelwelcomeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.StatusCode = 0
}

var poolTaobaoTmallgenieHotelwelcomeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmallgenieHotelwelcomeAPIResponse)
	},
}

// GetTaobaoTmallgenieHotelwelcomeAPIResponse 从 sync.Pool 获取 TaobaoTmallgenieHotelwelcomeAPIResponse
func GetTaobaoTmallgenieHotelwelcomeAPIResponse() *TaobaoTmallgenieHotelwelcomeAPIResponse {
	return poolTaobaoTmallgenieHotelwelcomeAPIResponse.Get().(*TaobaoTmallgenieHotelwelcomeAPIResponse)
}

// ReleaseTaobaoTmallgenieHotelwelcomeAPIResponse 将 TaobaoTmallgenieHotelwelcomeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmallgenieHotelwelcomeAPIResponse(v *TaobaoTmallgenieHotelwelcomeAPIResponse) {
	v.Reset()
	poolTaobaoTmallgenieHotelwelcomeAPIResponse.Put(v)
}
