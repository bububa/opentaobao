package car

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderQueryAPIResponse
飞猪订单状态查询接口 API返回值
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单 */
type TaobaoAlitripCarOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderQueryAPIResponseModel
}

// TaobaoAlitripCarOrderQueryAPIResponseModel is 飞猪订单状态查询接口 成功返回结果
type TaobaoAlitripCarOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单结果
	FirstResult *OrderQueryRsp `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
