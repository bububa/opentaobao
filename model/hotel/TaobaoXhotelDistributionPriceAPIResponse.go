package hotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDistributionPriceAPIResponse 飞猪分销通用酒店报价接口 API返回值
// taobao.xhotel.distribution.price
//
// 飞猪分销通用酒店报价接口
type TaobaoXhotelDistributionPriceAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelDistributionPriceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelDistributionPriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelDistributionPriceAPIResponseModel).Reset()
}

// TaobaoXhotelDistributionPriceAPIResponseModel is 飞猪分销通用酒店报价接口 成功返回结果
type TaobaoXhotelDistributionPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_distribution_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店报价查询结果
	HotelPricesResult *HotelPricesResult `json:"hotel_prices_result,omitempty" xml:"hotel_prices_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelDistributionPriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HotelPricesResult = nil
}

var poolTaobaoXhotelDistributionPriceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDistributionPriceAPIResponse)
	},
}

// GetTaobaoXhotelDistributionPriceAPIResponse 从 sync.Pool 获取 TaobaoXhotelDistributionPriceAPIResponse
func GetTaobaoXhotelDistributionPriceAPIResponse() *TaobaoXhotelDistributionPriceAPIResponse {
	return poolTaobaoXhotelDistributionPriceAPIResponse.Get().(*TaobaoXhotelDistributionPriceAPIResponse)
}

// ReleaseTaobaoXhotelDistributionPriceAPIResponse 将 TaobaoXhotelDistributionPriceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelDistributionPriceAPIResponse(v *TaobaoXhotelDistributionPriceAPIResponse) {
	v.Reset()
	poolTaobaoXhotelDistributionPriceAPIResponse.Put(v)
}
