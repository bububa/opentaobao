package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除购品 APIResponse
alibaba.wdk.marketing.fullrange.removeitem

删除换购商品
*/
type AlibabaWdkMarketingFullrangeRemoveitemAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingFullrangeRemoveitemResponse `json:"alibaba_wdk_marketing_fullrange_removeitem_response,omitempty"`
}

type AlibabaWdkMarketingFullrangeRemoveitemResponse struct {

    // result
    Result   *MarketResult `json:"result,omitempty"`

}
