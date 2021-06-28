package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部确认 APIResponse
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
*/
type TaobaoLifeReservationItemOrderConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLifeReservationItemOrderConfirmResponse `json:"life_reservation_item_order_confirm_response,omitempty"` 
    TaobaoLifeReservationItemOrderConfirmResponse
}

/* model for simplify = false
type TaobaoLifeReservationItemOrderConfirmResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoLifeReservationItemOrderConfirmResult  *TaobaoLifeReservationItemOrderConfirmResult `json:"taobao_life_reservation_item_order_confirm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoLifeReservationItemOrderConfirmResponse struct {

    // 接口返回model
    Result   *TaobaoLifeReservationItemOrderConfirmResult `json:"result,omitempty"`

}
