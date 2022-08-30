package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeBuyPlaceorderAPIResponse AE下单API API返回值
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
type AliexpressTradeBuyPlaceorderAPIResponse struct {
	model.CommonResponse
	AliexpressTradeBuyPlaceorderAPIResponseModel
}

// AliexpressTradeBuyPlaceorderAPIResponseModel is AE下单API 成功返回结果
type AliexpressTradeBuyPlaceorderAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_trade_buy_placeorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PlaceOrderRes4OpenApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
