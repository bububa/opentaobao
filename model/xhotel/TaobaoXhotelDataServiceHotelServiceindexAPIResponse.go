package xhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDataServiceHotelServiceindexAPIResponse 酒店服务指数 API返回值
// taobao.xhotel.data.service.hotel.serviceindex
//
// 酒店服务指数
type TaobaoXhotelDataServiceHotelServiceindexAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelDataServiceHotelServiceindexAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel).Reset()
}

// TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel is 酒店服务指数 成功返回结果
type TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_data_service_hotel_serviceindex_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelDataServiceHotelServiceindexResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelDataServiceHotelServiceindexAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDataServiceHotelServiceindexAPIResponse)
	},
}

// GetTaobaoXhotelDataServiceHotelServiceindexAPIResponse 从 sync.Pool 获取 TaobaoXhotelDataServiceHotelServiceindexAPIResponse
func GetTaobaoXhotelDataServiceHotelServiceindexAPIResponse() *TaobaoXhotelDataServiceHotelServiceindexAPIResponse {
	return poolTaobaoXhotelDataServiceHotelServiceindexAPIResponse.Get().(*TaobaoXhotelDataServiceHotelServiceindexAPIResponse)
}

// ReleaseTaobaoXhotelDataServiceHotelServiceindexAPIResponse 将 TaobaoXhotelDataServiceHotelServiceindexAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelDataServiceHotelServiceindexAPIResponse(v *TaobaoXhotelDataServiceHotelServiceindexAPIResponse) {
	v.Reset()
	poolTaobaoXhotelDataServiceHotelServiceindexAPIResponse.Put(v)
}
