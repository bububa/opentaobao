package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmallgenieHotelwelcomeAPIResponse
酒店欢迎词推送 API返回值
taobao.tmallgenie.hotelwelcome

推送欢迎词，让天猫精灵播放 */
type TaobaoTmallgenieHotelwelcomeAPIResponse struct {
	model.CommonResponse
	TaobaoTmallgenieHotelwelcomeAPIResponseModel
}

// TaobaoTmallgenieHotelwelcomeAPIResponseModel is 酒店欢迎词推送 成功返回结果
type TaobaoTmallgenieHotelwelcomeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmallgenie_hotelwelcome_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
