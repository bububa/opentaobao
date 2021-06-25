package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
促销价签服务 APIResponse
alibaba.wdk.marketing.price

获取营销-促销商品中的实时价格
*/
type AlibabaWdkMarketingPriceAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingPriceResponse `json:"alibaba_wdk_marketing_price_response,omitempty"`
}

type AlibabaWdkMarketingPriceResponse struct {

    // 返回结果
    Result   *PromotionPriceResult `json:"result,omitempty"`

}
