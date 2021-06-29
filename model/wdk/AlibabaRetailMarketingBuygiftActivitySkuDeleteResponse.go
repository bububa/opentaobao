package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动商品 API返回值 
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingBuygiftActivitySkuDeleteResponse
}

// 删除单品买赠活动商品 成功返回结果
type AlibabaRetailMarketingBuygiftActivitySkuDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_sku_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
