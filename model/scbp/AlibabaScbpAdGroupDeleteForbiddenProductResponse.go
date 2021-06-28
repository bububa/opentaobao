package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除屏蔽品 APIResponse
alibaba.scbp.ad.group.delete.forbidden.product

删除屏蔽品
*/
type AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupDeleteForbiddenProductResponse
}

type AlibabaScbpAdGroupDeleteForbiddenProductResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_delete_forbidden_product_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
