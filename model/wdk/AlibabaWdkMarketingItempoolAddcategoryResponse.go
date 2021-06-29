package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品池里面的类目 API返回值 
alibaba.wdk.marketing.itempool.addcategory

增加商品池里面的类目
*/
type AlibabaWdkMarketingItempoolAddcategoryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolAddcategoryResponse
}

// 增加商品池里面的类目 成功返回结果
type AlibabaWdkMarketingItempoolAddcategoryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_addcategory_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
