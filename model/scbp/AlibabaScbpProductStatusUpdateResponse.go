package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改P4P产品推广状态 API返回值 
alibaba.scbp.product.status.update

修改P4P产品推广状态
*/
type AlibabaScbpProductStatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpProductStatusUpdateResponse
}

// 修改P4P产品推广状态 成功返回结果
type AlibabaScbpProductStatusUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_product_status_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 实际修改的产品数
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
