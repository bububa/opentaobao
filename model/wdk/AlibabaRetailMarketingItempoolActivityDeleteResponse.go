package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动【同城零售】 API返回值 
alibaba.retail.marketing.itempool.activity.delete

同城零售商品池活动删除
*/
type AlibabaRetailMarketingItempoolActivityDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingItempoolActivityDeleteResponse
}

// 删除商品池活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolActivityDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
