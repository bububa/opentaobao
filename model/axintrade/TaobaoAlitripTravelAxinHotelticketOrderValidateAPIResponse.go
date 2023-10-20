package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketordervalidateAPIResponse 阿信度假业务交易试单接口 API返回值
// taobao.alitrip.travel.axin.hotelticket.order.validate
//
// 阿信度假业务交易试单接口
type TaobaoalitriptravelaxinhotelticketordervalidateAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelticketordervalidateAPIResponseModel
}

// TaobaoalitriptravelaxinhotelticketordervalidateAPIResponseModel is 阿信度假业务交易试单接口 成功返回结果
type TaobaoalitriptravelaxinhotelticketordervalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_order_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回类
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
