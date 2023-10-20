package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticssellerordersgetAPIResponse 商家配送核销订单列表查询 API返回值
// alibaba.ascp.logistics.seller.orders.get
//
// 商家配送核销订单列表查询
type AlibabaascplogisticssellerordersgetAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticssellerordersgetAPIResponseModel
}

// AlibabaascplogisticssellerordersgetAPIResponseModel is 商家配送核销订单列表查询 成功返回结果
type AlibabaascplogisticssellerordersgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
