package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketProductListAPIResponse 阿信酒景套餐产品列表查询 API返回值
// taobao.alitrip.travel.axin.hotelticket.product.list
//
// 阿信酒景套餐产品列表查询
type TaobaoAlitripTravelAxinHotelticketProductListAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketProductListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketProductListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinHotelticketProductListAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinHotelticketProductListAPIResponseModel is 阿信酒景套餐产品列表查询 成功返回结果
type TaobaoAlitripTravelAxinHotelticketProductListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinHotelticketProductListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinHotelticketProductListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelticketProductListAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinHotelticketProductListAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketProductListAPIResponse
func GetTaobaoAlitripTravelAxinHotelticketProductListAPIResponse() *TaobaoAlitripTravelAxinHotelticketProductListAPIResponse {
	return poolTaobaoAlitripTravelAxinHotelticketProductListAPIResponse.Get().(*TaobaoAlitripTravelAxinHotelticketProductListAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketProductListAPIResponse 将 TaobaoAlitripTravelAxinHotelticketProductListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketProductListAPIResponse(v *TaobaoAlitripTravelAxinHotelticketProductListAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketProductListAPIResponse.Put(v)
}
