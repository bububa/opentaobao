package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
司机应答API APIResponse
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
*/
type TaobaoAlitripCarOrderConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_car_order_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"