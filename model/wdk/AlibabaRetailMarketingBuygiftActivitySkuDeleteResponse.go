package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动商品 APIResponse
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingBuygiftActivitySkuDeleteResponse
}

type AlibabaRetailMarketingBuygiftActivitySkuDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_sku_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
