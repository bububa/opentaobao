package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品池活动【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.update

同城零售商品池活动更新
*/
type AlibabaRetailMarketingItempoolActivityUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingItempoolActivityUpdateResponse
}

type AlibabaRetailMarketingItempoolActivityUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
