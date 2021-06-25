package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品池活动下的商品 APIResponse
alibaba.wdk.marketing.itempool.queryitems

查询商品池活动下面的商品
*/
type AlibabaWdkMarketingItempoolQueryitemsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolQueryitemsResponse `json:"alibaba_wdk_marketing_itempool_queryitems_response,omitempty"`
}

type AlibabaWdkMarketingItempoolQueryitemsResponse struct {

    // 查询返回结果
    Result   *MarketPageResult `json:"result,omitempty"`

}
