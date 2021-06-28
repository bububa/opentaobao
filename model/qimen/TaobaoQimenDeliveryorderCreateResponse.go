package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建接口 APIResponse
taobao.qimen.deliveryorder.create

taobao.qimen.deliveryorder.create
*/
type TaobaoQimenDeliveryorderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenDeliveryorderCreateResponse
}

type TaobaoQimenDeliveryorderCreateResponse struct {
    XMLName xml.Name `xml:"qimen_deliveryorder_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *DeliveryOrderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`

    
}
