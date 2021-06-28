package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪订单状态查询接口 APIResponse
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单
*/
type TaobaoAlitripCarOrderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_car_order_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单结果
    
    FirstResult   *OrderQueryRsp `json:"first_result,omitempty" xml:"