package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新版跟团游商品维护接口 APIResponse
alitrip.grouptour.product.upload

新版跟团游商品维护接口
*/
type AlitripGrouptourProductUploadAPIResponse struct {
    model.CommonResponse
    Response *AlitripGrouptourProductUploadResponse `json:"alitrip_grouptour_product_upload_response,omitempty"`
}

type AlitripGrouptourProductUploadResponse struct {

    // firstResult
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
