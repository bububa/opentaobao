package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
全场增加换购品 APIResponse
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品
*/
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingFullrangeAddexchangeitemResponse `json:"alibaba_wdk_marketing_fullrange_addexchangeitem_response,omitempty"`
}

type AlibabaWdkMarketingFullrangeAddexchangeitemResponse struct {

    // 添加商品返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
