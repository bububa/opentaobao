package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscHotelListQueryAPIResponse 标准酒店信息查询-供销平台 API返回值
// taobao.alitrip.travel.fsc.hotel.list.query
//
// 供销平台标准酒店信息列表查询
type TaobaoAlitripTravelFscHotelListQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscHotelListQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscHotelListQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscHotelListQueryAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscHotelListQueryAPIResponseModel is 标准酒店信息查询-供销平台 成功返回结果
type TaobaoAlitripTravelFscHotelListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_hotel_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口应答对象
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscHotelListQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelFscHotelListQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscHotelListQueryAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscHotelListQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscHotelListQueryAPIResponse
func GetTaobaoAlitripTravelFscHotelListQueryAPIResponse() *TaobaoAlitripTravelFscHotelListQueryAPIResponse {
	return poolTaobaoAlitripTravelFscHotelListQueryAPIResponse.Get().(*TaobaoAlitripTravelFscHotelListQueryAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscHotelListQueryAPIResponse 将 TaobaoAlitripTravelFscHotelListQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscHotelListQueryAPIResponse(v *TaobaoAlitripTravelFscHotelListQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscHotelListQueryAPIResponse.Put(v)
}
