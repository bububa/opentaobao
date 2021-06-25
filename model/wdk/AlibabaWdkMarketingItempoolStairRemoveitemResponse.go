package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除换购活动商品 APIResponse
alibaba.wdk.marketing.itempool.stair.removeitem

删除换购商品
*/
type AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolStairRemoveitemResponse `json:"alibaba_wdk_marketing_itempool_stair_removeitem_response,omitempty"`
}

type AlibabaWdkMarketingItempoolStairRemoveitemResponse struct {

    // result
    Result   *MarketResult `json:"result,omitempty"`

}
