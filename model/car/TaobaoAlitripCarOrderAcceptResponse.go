package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认接单 APIResponse
taobao.alitrip.car.order.accept

用来接收服务商确认接单信息
*/
type TaobaoAlitripCarOrderAcceptAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripCarOrderAcceptResponse
}

type TaobaoAlitripCarOrderAcceptResponse struct {
    XMLName xml.Name `xml:"alitrip_car_order_accept_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 根据站点名称查询产品
    
    Result   *TaobaoAlitripCarOrderAcceptApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
