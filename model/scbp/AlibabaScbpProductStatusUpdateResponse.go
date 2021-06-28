package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改P4P产品推广状态 APIResponse
alibaba.scbp.product.status.update

修改P4P产品推广状态
*/
type AlibabaScbpProductStatusUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpProductStatusUpdateResponse `json:"alibaba_scbp_product_status_update_response,omitempty"` 
    AlibabaScbpProductStatusUpdateResponse
}

/* model for simplify = false
type AlibabaScbpProductStatusUpdateResponse struct {

    // 实际修改的产品数
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpProductStatusUpdateResponse struct {

    // 实际修改的产品数
    Result   int64 `json:"result,omitempty"`

}
