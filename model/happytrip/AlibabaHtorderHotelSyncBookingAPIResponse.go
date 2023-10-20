package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHtorderHotelSyncBookingAPIResponse 未来酒店亲橙客栈预订信息同步 API返回值
// alibaba.htorder.hotel.sync.booking
//
// 未来酒店亲橙客栈预订信息同步
type AlibabaHtorderHotelSyncBookingAPIResponse struct {
	model.CommonResponse
	AlibabaHtorderHotelSyncBookingAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHtorderHotelSyncBookingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHtorderHotelSyncBookingAPIResponseModel).Reset()
}

// AlibabaHtorderHotelSyncBookingAPIResponseModel is 未来酒店亲橙客栈预订信息同步 成功返回结果
type AlibabaHtorderHotelSyncBookingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_htorder_hotel_sync_booking_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaHtorderHotelSyncBookingResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHtorderHotelSyncBookingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHtorderHotelSyncBookingAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHtorderHotelSyncBookingAPIResponse)
	},
}

// GetAlibabaHtorderHotelSyncBookingAPIResponse 从 sync.Pool 获取 AlibabaHtorderHotelSyncBookingAPIResponse
func GetAlibabaHtorderHotelSyncBookingAPIResponse() *AlibabaHtorderHotelSyncBookingAPIResponse {
	return poolAlibabaHtorderHotelSyncBookingAPIResponse.Get().(*AlibabaHtorderHotelSyncBookingAPIResponse)
}

// ReleaseAlibabaHtorderHotelSyncBookingAPIResponse 将 AlibabaHtorderHotelSyncBookingAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHtorderHotelSyncBookingAPIResponse(v *AlibabaHtorderHotelSyncBookingAPIResponse) {
	v.Reset()
	poolAlibabaHtorderHotelSyncBookingAPIResponse.Put(v)
}
