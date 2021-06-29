package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
版本数量查询 APIResponse
alibaba.wdk.marketing.open.version.count

版本数量查询
*/
type AlibabaWdkMarketingOpenVersionCountAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingOpenVersionCountResponse
}

type AlibabaWdkMarketingOpenVersionCountResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_open_version_count_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
