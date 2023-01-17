package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleCarOrderQueryAPIResponse 二手车寄卖查询订单接口 API返回值
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
type AlibabaIdleCarOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleCarOrderQueryAPIResponseModel
}

// AlibabaIdleCarOrderQueryAPIResponseModel is 二手车寄卖查询订单接口 成功返回结果
type AlibabaIdleCarOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_car_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaIdleCarOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
