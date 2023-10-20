package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponse 阿信度假业务申请退款 API返回值
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
type TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponseModel
}

// TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponseModel is 阿信度假业务申请退款 成功返回结果
type TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_refund_orderrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回类
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
