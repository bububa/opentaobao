package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新
*/
type AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingBuygiftActivityUpdateResponse
}

type AlibabaRetailMarketingBuygiftActivityUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
