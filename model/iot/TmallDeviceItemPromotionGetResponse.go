package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上商品优惠获取 APIResponse
tmall.device.item.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TmallDeviceItemPromotionGetAPIResponse struct {
    model.CommonResponse
    Response *TmallDeviceItemPromotionGetResponse `json:"tmall_device_item_promotion_get_response,omitempty"`
}

type TmallDeviceItemPromotionGetResponse struct {

    // 优惠详细信息
    Promotions   *PromotionDisplayTop `json:"promotions,omitempty"`

}
