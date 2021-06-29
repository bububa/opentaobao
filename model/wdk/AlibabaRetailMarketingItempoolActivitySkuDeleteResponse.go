package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动商品【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.sku.delete

删除商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingItempoolActivitySkuDeleteResponse
}

type AlibabaRetailMarketingItempoolActivitySkuDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_sku_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
