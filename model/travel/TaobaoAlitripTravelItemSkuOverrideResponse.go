package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】商品级别日历价格库存修改，全量覆盖 APIResponse
taobao.alitrip.travel.item.sku.override

旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
*/
type TaobaoAlitripTravelItemSkuOverrideAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_item_sku_override_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品sku修改结果
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"