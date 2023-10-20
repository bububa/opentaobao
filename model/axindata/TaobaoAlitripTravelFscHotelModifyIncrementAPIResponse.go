package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse 酒店价库变更列表查询-供销平台 API返回值
// taobao.alitrip.travel.fsc.hotel.modify.increment
//
// 按照时间纬度查询酒店变更列表
type TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel).Reset()
}

// TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel is 酒店价库变更列表查询-供销平台 成功返回结果
type TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_hotel_modify_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口应答对象
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelFscHotelModifyIncrementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse)
	},
}

// GetTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse
func GetTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse() *TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse {
	return poolTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse.Get().(*TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse)
}

// ReleaseTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse 将 TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse(v *TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelFscHotelModifyIncrementAPIResponse.Put(v)
}
