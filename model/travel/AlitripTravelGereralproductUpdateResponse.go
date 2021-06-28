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
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_gereralproduct_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // firstResult
    
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"