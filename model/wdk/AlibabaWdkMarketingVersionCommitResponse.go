package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交版本号 APIResponse
alibaba.wdk.marketing.version.commit

提交版本号，标识结束此版本操作
*/
type AlibabaWdkMarketingVersionCommitAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingVersionCommitResponse
}

type AlibabaWdkMarketingVersionCommitResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_version_commit_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
