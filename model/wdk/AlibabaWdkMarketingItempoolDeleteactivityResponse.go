package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动 APIResponse
alibaba.wdk.marketing.itempool.deleteactivity

删除商品池活动
*/
type AlibabaWdkMarketingItempoolDeleteactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolDeleteactivityResponse `json:"alibaba_wdk_marketing_itempool_deleteactivity_response,omitempty"`
}

type AlibabaWdkMarketingItempoolDeleteactivityResponse struct {

    // 删除活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
