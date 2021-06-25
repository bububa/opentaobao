package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新单品买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新
*/
type AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingBuygiftActivityUpdateResponse `json:"alibaba_retail_marketing_buygift_activity_update_response,omitempty"`
}

type AlibabaRetailMarketingBuygiftActivityUpdateResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
