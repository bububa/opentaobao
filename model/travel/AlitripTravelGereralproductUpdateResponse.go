package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通用类目产品发布编辑 APIResponse
alitrip.travel.gereralproduct.update

提供给飞猪供销平台供应商发布编辑通用类目产品的API
*/
type AlitripTravelGereralproductUpdateAPIResponse struct {
    model.CommonResponse
    AlitripTravelGereralproductUpdateResponse
}

type AlitripTravelGereralproductUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_gereralproduct_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // firstResult
    
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
