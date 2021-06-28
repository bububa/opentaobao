package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改P4P产品推广状态 APIResponse
alibaba.scbp.product.status.update

修改P4P产品推广状态
*/
type AlibabaScbpProductStatusUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_product_status_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 实际修改的产品数
    
    Result   int64 `json:"result,omitempty" xml:"