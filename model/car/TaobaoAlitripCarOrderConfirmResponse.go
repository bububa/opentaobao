package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
司机应答API API返回值 
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
*/
type TaobaoAlitripCarOrderConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripCarOrderConfirmResponse
}

// 司机应答API 成功返回结果
type TaobaoAlitripCarOrderConfirmResponse struct {
    XMLName xml.Name `xml:"alitrip_car_order_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
    // 其它数据
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
