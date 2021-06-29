package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽品 APIResponse
alibaba.scbp.ad.group.create.forbidden.product

创建屏蔽品
*/
type AlibabaScbpAdGroupCreateForbiddenProductAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupCreateForbiddenProductResponse
}

type AlibabaScbpAdGroupCreateForbiddenProductResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_create_forbidden_product_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
