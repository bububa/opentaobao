package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询货架和位置数据 APIResponse
alibaba.interact.retail.queryshelflocation

查询货架和位置数据
*/
type AlibabaInteractRetailQueryshelflocationAPIResponse struct {
    model.CommonResponse
    AlibabaInteractRetailQueryshelflocationResponse
}

type AlibabaInteractRetailQueryshelflocationResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_retail_queryshelflocation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaInteractRetailQueryshelflocationResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
