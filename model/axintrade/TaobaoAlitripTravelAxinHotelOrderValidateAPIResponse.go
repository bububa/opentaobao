package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelordervalidateAPIResponse 阿信酒店订单数据校验 API返回值
// taobao.alitrip.travel.axin.hotel.order.validate
//
// 阿信酒店订单下单前的数据校验，包括酒店、房型、售卖政策、入离日期等参数的校验
type TaobaoalitriptravelaxinhotelordervalidateAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelordervalidateAPIResponseModel
}

// TaobaoalitriptravelaxinhotelordervalidateAPIResponseModel is 阿信酒店订单数据校验 成功返回结果
type TaobaoalitriptravelaxinhotelordervalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
