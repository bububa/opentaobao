package gameact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户收件地址确认 API返回值 
taobao.de.activity.delivery.addr.confirm

用户收件地址确认
*/
type TaobaoDeActivityDeliveryAddrConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoDeActivityDeliveryAddrConfirmResponse
}

// 用户收件地址确认 成功返回结果
type TaobaoDeActivityDeliveryAddrConfirmResponse struct {
    XMLName xml.Name `xml:"de_activity_delivery_addr_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 更新或确认收件地址
    UpdateDeliveryAddressVo   *UpdateDeliveryAddressVO `json:"update_delivery_address_vo,omitempty" xml:"update_delivery_address_vo,omitempty"`
}
