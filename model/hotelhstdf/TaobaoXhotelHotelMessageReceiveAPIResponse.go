package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelHotelMessageReceiveAPIResponse 接收道消息接口 API返回值
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
type TaobaoXhotelHotelMessageReceiveAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelHotelMessageReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelHotelMessageReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelHotelMessageReceiveAPIResponseModel).Reset()
}

// TaobaoXhotelHotelMessageReceiveAPIResponseModel is 接收道消息接口 成功返回结果
type TaobaoXhotelHotelMessageReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_hotel_message_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *DidaResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelHotelMessageReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelHotelMessageReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelHotelMessageReceiveAPIResponse)
	},
}

// GetTaobaoXhotelHotelMessageReceiveAPIResponse 从 sync.Pool 获取 TaobaoXhotelHotelMessageReceiveAPIResponse
func GetTaobaoXhotelHotelMessageReceiveAPIResponse() *TaobaoXhotelHotelMessageReceiveAPIResponse {
	return poolTaobaoXhotelHotelMessageReceiveAPIResponse.Get().(*TaobaoXhotelHotelMessageReceiveAPIResponse)
}

// ReleaseTaobaoXhotelHotelMessageReceiveAPIResponse 将 TaobaoXhotelHotelMessageReceiveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelHotelMessageReceiveAPIResponse(v *TaobaoXhotelHotelMessageReceiveAPIResponse) {
	v.Reset()
	poolTaobaoXhotelHotelMessageReceiveAPIResponse.Put(v)
}
