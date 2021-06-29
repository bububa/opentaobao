package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部发起变更 API返回值 
taobao.life.reservation.item.order.change

生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力
*/
type TaobaoLifeReservationItemOrderChangeAPIResponse struct {
    model.CommonResponse
    TaobaoLifeReservationItemOrderChangeResponse
}

// 生服购后预约单外部发起变更 成功返回结果
type TaobaoLifeReservationItemOrderChangeResponse struct {
    XMLName xml.Name `xml:"life_reservation_item_order_change_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoLifeReservationItemOrderChangeResult `json:"result,omitempty" xml:"result,omitempty"`
}
