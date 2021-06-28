package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设置P4P产品优先推广状态 APIResponse
alibaba.scbp.product.preferential.update

设置P4P产品优先推广状态
*/
type AlibabaScbpProductPreferentialUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpProductPreferentialUpdateResponse `json:"alibaba_scbp_product_preferential_update_response,omitempty"` 
    AlibabaScbpProductPreferentialUpdateResponse
}

/* model for simplify = false
type AlibabaScbpProductPreferentialUpdateResponse struct {

    // 设置成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaScbpProductPreferentialUpdateResponse struct {

    // 设置成功
    Result   bool `json:"result,omitempty"`

}
