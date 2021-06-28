package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提交版本号 APIResponse
alibaba.wdk.marketing.version.commit

提交版本号，标识结束此版本操作
*/
type AlibabaWdkMarketingVersionCommitAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingVersionCommitResponse `json:"alibaba_wdk_marketing_version_commit_response,omitempty"` 
    AlibabaWdkMarketingVersionCommitResponse
}

/* model for simplify = false
type AlibabaWdkMarketingVersionCommitResponse struct {

    // result
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingVersionCommitResponse struct {

    // result
    Result   *MarketResult `json:"result,omitempty"`

}
