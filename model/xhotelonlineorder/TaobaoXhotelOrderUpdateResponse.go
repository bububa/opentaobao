package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店订单发货接口 APIResponse
taobao.xhotel.order.update

卖家确认订单或者取消订单，适用于预付、面付、信用住订单
*/
type TaobaoXhotelOrderUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderUpdateResponse
}

type TaobaoXhotelOrderUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_order_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回提示信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
