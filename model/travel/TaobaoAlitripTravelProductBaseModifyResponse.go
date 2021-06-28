package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商编辑产品API APIResponse
taobao.alitrip.travel.product.base.modify

飞猪供销平台供应商可通过该API编辑产品
*/
type TaobaoAlitripTravelProductBaseModifyAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelProductBaseModifyResponse
}

type TaobaoAlitripTravelProductBaseModifyResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_product_base_modify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品修改结果
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`

    
}
