package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部确认 APIResponse
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
*/
type TaobaoLifeReservationItemOrderConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"life_reservation_item_order_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoLifeReservationItemOrderConfirmResult `json:"result,omitempty" xml:"