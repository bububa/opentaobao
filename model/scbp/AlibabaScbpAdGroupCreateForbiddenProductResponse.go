package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽品 APIResponse
alibaba.scbp.ad.group.create.forbidden.product

创建屏蔽品
*/
type AlibabaScbpAdGroupCreateForbiddenProductAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdGroupCreateForbiddenProductResponse `json:"alibaba_scbp_ad_group_create_forbidden_product_response,omitempty"` 
    AlibabaScbpAdGroupCreateForbiddenProductResponse
}

/* model for simplify = false
type AlibabaScbpAdGroupCreateForbiddenProductResponse struct {

    // 返回结果
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdGroupCreateForbiddenProductResponse struct {

    // 返回结果
    Result   int64 `json:"result,omitempty"`

}
