package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除屏蔽品 APIResponse
alibaba.scbp.ad.group.delete.forbidden.product

删除屏蔽品
*/
type AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdGroupDeleteForbiddenProductResponse `json:"alibaba_scbp_ad_group_delete_forbidden_product_response,omitempty"` 
    AlibabaScbpAdGroupDeleteForbiddenProductResponse
}

/* model for simplify = false
type AlibabaScbpAdGroupDeleteForbiddenProductResponse struct {

    // 返回值
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdGroupDeleteForbiddenProductResponse struct {

    // 返回值
    Result   int64 `json:"result,omitempty"`

}
