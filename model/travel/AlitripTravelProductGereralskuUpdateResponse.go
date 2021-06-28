package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(供销)船票通用类目sku新增&编辑API APIResponse
alitrip.travel.product.gereralsku.update

发布SKU信息（如果properties重复 则更新）
*/
type AlitripTravelProductGereralskuUpdateAPIResponse struct {
    model.CommonResponse
    AlitripTravelProductGereralskuUpdateResponse
}

type AlitripTravelProductGereralskuUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_product_gereralsku_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
