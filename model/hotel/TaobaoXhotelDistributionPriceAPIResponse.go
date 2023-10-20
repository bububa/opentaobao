package hotel

import (
	"encoding/xml"

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

// TaobaoXhotelDistributionPriceAPIResponseModel is 飞猪分销通用酒店报价接口 成功返回结果
type TaobaoXhotelDistributionPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_distribution_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店报价查询结果
	HotelPricesResult *HotelPricesResult `json:"hotel_prices_result,omitempty" xml:"hotel_prices_result,omitempty"`
}
