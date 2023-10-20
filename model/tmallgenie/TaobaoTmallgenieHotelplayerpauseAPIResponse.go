package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmallgenieHotelplayerpauseAPIResponse 天猫精灵酒店播放暂停 API返回值
// taobao.tmallgenie.hotelplayerpause
//
// 酒店推送指令给天猫精灵停止播放音乐
type TaobaoTmallgenieHotelplayerpauseAPIResponse struct {
	model.CommonResponse
	TaobaoTmallgenieHotelplayerpauseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmallgenieHotelplayerpauseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmallgenieHotelplayerpauseAPIResponseModel).Reset()
}

// TaobaoTmallgenieHotelplayerpauseAPIResponseModel is 天猫精灵酒店播放暂停 成功返回结果
type TaobaoTmallgenieHotelplayerpauseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmallgenie_hotelplayerpause_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmallgenieHotelplayerpauseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.StatusCode = 0
}

var poolTaobaoTmallgenieHotelplayerpauseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmallgenieHotelplayerpauseAPIResponse)
	},
}

// GetTaobaoTmallgenieHotelplayerpauseAPIResponse 从 sync.Pool 获取 TaobaoTmallgenieHotelplayerpauseAPIResponse
func GetTaobaoTmallgenieHotelplayerpauseAPIResponse() *TaobaoTmallgenieHotelplayerpauseAPIResponse {
	return poolTaobaoTmallgenieHotelplayerpauseAPIResponse.Get().(*TaobaoTmallgenieHotelplayerpauseAPIResponse)
}

// ReleaseTaobaoTmallgenieHotelplayerpauseAPIResponse 将 TaobaoTmallgenieHotelplayerpauseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmallgenieHotelplayerpauseAPIResponse(v *TaobaoTmallgenieHotelplayerpauseAPIResponse) {
	v.Reset()
	poolTaobaoTmallgenieHotelplayerpauseAPIResponse.Put(v)
}
