package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelordercreateAPIResponse 酒店分销订单创建服务-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
type TaobaoalitriptravelaxinhotelordercreateAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelordercreateAPIResponseModel
}

// TaobaoalitriptravelaxinhotelordercreateAPIResponseModel is 酒店分销订单创建服务-阿信 成功返回结果
type TaobaoalitriptravelaxinhotelordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitriptravelaxinhotelordercreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
