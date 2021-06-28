package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商新增产品API APIResponse
taobao.alitrip.travel.product.base.add

飞猪供销平台供应商可通过该API发布新产品
*/
type TaobaoAlitripTravelProductBaseAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_product_base_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品发布结果
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"