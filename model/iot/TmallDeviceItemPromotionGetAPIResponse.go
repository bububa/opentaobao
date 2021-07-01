package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上商品优惠获取 API返回值 
tmall.device.item.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TmallDeviceItemPromotionGetAPIResponse struct {
    model.CommonResponse
    TmallDeviceItemPromotionGetAPIResponseModel
}

// 智能硬件上商品优惠获取 成功返回结果
type TmallDeviceItemPromotionGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_device_item_promotion_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 优惠详细信息
    Promotions   *PromotionDisplayTop `json:"promotions,omitempty" xml:"promotions,omitempty"`
}
