package gameact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户收件地址确认 APIResponse
taobao.de.activity.delivery.addr.confirm

用户收件地址确认
*/
type TaobaoDeActivityDeliveryAddrConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoDeActivityDeliveryAddrConfirmResponse
}

type TaobaoDeActivityDeliveryAddrConfirmResponse struct {
    XMLName xml.Name `xml:"de_activity_delivery_addr_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新或确认收件地址
    
    UpdateDeliveryAddressVo   *UpdateDeliveryAddressVO `json:"update_delivery_address_vo,omitempty" xml:"update_delivery_address_vo,omitempty"`

    
}
