package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】日期级别日历价格库存修改，增量维护 APIResponse
taobao.alitrip.travel.item.sku.price.modify

【API3.0】日期级别日历价格库存增量维护
*/
type TaobaoAlitripTravelItemSkuPriceModifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_item_sku_price_modify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 日期级别日历价格库存增量维护
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"