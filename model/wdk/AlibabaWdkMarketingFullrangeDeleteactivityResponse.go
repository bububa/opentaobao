package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除活动接口 APIResponse
alibaba.wdk.marketing.fullrange.deleteactivity

全场活动删除活动
*/
type AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingFullrangeDeleteactivityResponse `json:"alibaba_wdk_marketing_fullrange_deleteactivity_response,omitempty"`
}

type AlibabaWdkMarketingFullrangeDeleteactivityResponse struct {

    // 删除活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
