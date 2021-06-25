package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingBuygiftActivityDeleteResponse `json:"alibaba_retail_marketing_buygift_activity_delete_response,omitempty"`
}

type AlibabaRetailMarketingBuygiftActivityDeleteResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
