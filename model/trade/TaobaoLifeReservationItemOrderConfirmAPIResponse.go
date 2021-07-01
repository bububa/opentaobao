package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部确认 API返回值 
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
*/
type TaobaoLifeReservationItemOrderConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoLifeReservationItemOrderConfirmAPIResponseModel
}

// 生服购后预约单外部确认 成功返回结果
type TaobaoLifeReservationItemOrderConfirmAPIResponseModel struct {
    XMLName xml.Name `xml:"life_reservation_item_order_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoLifeReservationItemOrderConfirmResult `json:"result,omitempty" xml:"result,omitempty"`
}
