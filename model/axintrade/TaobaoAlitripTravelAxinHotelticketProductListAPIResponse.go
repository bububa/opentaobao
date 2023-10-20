package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketproductlistAPIResponse 阿信酒景套餐产品列表查询 API返回值
// taobao.alitrip.travel.axin.hotelticket.product.list
//
// 阿信酒景套餐产品列表查询
type TaobaoalitriptravelaxinhotelticketproductlistAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelticketproductlistAPIResponseModel
}

// TaobaoalitriptravelaxinhotelticketproductlistAPIResponseModel is 阿信酒景套餐产品列表查询 成功返回结果
type TaobaoalitriptravelaxinhotelticketproductlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
