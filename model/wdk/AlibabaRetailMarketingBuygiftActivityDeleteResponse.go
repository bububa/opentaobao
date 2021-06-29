package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动 API返回值 
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaRetailMarketingBuygiftActivityDeleteResponse
}

// 删除单品买赠活动 成功返回结果
type AlibabaRetailMarketingBuygiftActivityDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
