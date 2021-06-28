package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票订单查询 APIResponse
taobao.bus.order.get

商家汽车票订单查询
*/
type TaobaoBusOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusOrderGetResponse
}

type TaobaoBusOrderGetResponse struct {
    XMLName xml.Name `xml:"bus_order_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单查询返回对象
    
    Result   *B2BOrderQueryRp `json:"result,omitempty" xml:"result,omitempty"`

    
}
