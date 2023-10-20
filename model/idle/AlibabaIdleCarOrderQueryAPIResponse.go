package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlecarorderqueryAPIResponse 二手车寄卖查询订单接口 API返回值
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
type AlibabaidlecarorderqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidlecarorderqueryAPIResponseModel
}

// AlibabaidlecarorderqueryAPIResponseModel is 二手车寄卖查询订单接口 成功返回结果
type AlibabaidlecarorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_car_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaidlecarorderqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
