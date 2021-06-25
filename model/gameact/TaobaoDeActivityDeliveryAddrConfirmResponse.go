package gameact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户收件地址确认 APIResponse
taobao.de.activity.delivery.addr.confirm

用户收件地址确认
*/
type TaobaoDeActivityDeliveryAddrConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeActivityDeliveryAddrConfirmResponse `json:"taobao_de_activity_delivery_addr_confirm_response,omitempty"`
}

type TaobaoDeActivityDeliveryAddrConfirmResponse struct {

    // 更新或确认收件地址
    UpdateDeliveryAddressVo   *UpdateDeliveryAddressVO `json:"update_delivery_address_vo,omitempty"`

}
