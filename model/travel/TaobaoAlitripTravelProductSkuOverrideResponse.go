package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（供销）产品级别日历价格库存修改，全量覆盖 APIResponse
taobao.alitrip.travel.product.sku.override

（供销）产品级别日历价格库存修改，全量覆盖
*/
type TaobaoAlitripTravelProductSkuOverrideAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_product_sku_override_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品sku修改结果
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"