package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新版跟团游商品维护接口 API返回值 
alitrip.grouptour.product.upload

新版跟团游商品维护接口
*/
type AlitripGrouptourProductUploadAPIResponse struct {
    model.CommonResponse
    AlitripGrouptourProductUploadAPIResponseModel
}

// 新版跟团游商品维护接口 成功返回结果
type AlitripGrouptourProductUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_grouptour_product_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // firstResult
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
