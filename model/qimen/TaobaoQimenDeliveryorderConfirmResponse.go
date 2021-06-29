package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 APIResponse
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm
*/
type TaobaoQimenDeliveryorderConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoQimenDeliveryorderConfirmResponse
}

type TaobaoQimenDeliveryorderConfirmResponse struct {
    XMLName xml.Name `xml:"qimen_deliveryorder_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
