package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加单品买赠活动商品 APIResponse
alibaba.retail.marketing.buygift.activity.sku.add

新增或更新单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingBuygiftActivitySkuAddResponse
}

type AlibabaRetailMarketingBuygiftActivitySkuAddResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_sku_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
