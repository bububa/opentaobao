package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.create

同城供给买赠活动创建
*/
type AlibabaRetailMarketingBuygiftActivityCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingBuygiftActivityCreateResponse `json:"alibaba_retail_marketing_buygift_activity_create_response,omitempty"`
}

type AlibabaRetailMarketingBuygiftActivityCreateResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}