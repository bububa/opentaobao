package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询活动 APIResponse
alibaba.wdk.marketing.fullrange.queryactivity

全场活动查询活动
*/
type AlibabaWdkMarketingFullrangeQueryactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingFullrangeQueryactivityResponse `json:"alibaba_wdk_marketing_fullrange_queryactivity_response,omitempty"`
}

type AlibabaWdkMarketingFullrangeQueryactivityResponse struct {

    // 查询返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
