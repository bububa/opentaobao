package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除活动接口 APIResponse
alibaba.wdk.marketing.fullrange.deleteactivity

全场活动删除活动
*/
type AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingFullrangeDeleteactivityResponse
}

type AlibabaWdkMarketingFullrangeDeleteactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_deleteactivity_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
