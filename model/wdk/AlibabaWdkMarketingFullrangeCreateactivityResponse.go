package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建全场活动 APIResponse
alibaba.wdk.marketing.fullrange.createactivity

创建全场活动
*/
type AlibabaWdkMarketingFullrangeCreateactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingFullrangeCreateactivityResponse `json:"alibaba_wdk_marketing_fullrange_createactivity_response,omitempty"`
}

type AlibabaWdkMarketingFullrangeCreateactivityResponse struct {

    // 创建活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
