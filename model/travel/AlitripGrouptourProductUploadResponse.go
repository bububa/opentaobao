package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新版跟团游商品维护接口 APIResponse
alitrip.grouptour.product.upload

新版跟团游商品维护接口
*/
type AlitripGrouptourProductUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_grouptour_product_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // firstResult
    
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"