package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的类目 APIResponse
alibaba.wdk.marketing.itempool.addcategory

增加商品池里面的类目
*/
type AlibabaWdkMarketingItempoolAddcategoryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolAddcategoryResponse `json:"alibaba_wdk_marketing_itempool_addcategory_response,omitempty"`
}

type AlibabaWdkMarketingItempoolAddcategoryResponse struct {

    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
