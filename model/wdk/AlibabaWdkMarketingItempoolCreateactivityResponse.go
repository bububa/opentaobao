package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品池活动 APIResponse
alibaba.wdk.marketing.itempool.createactivity

添加商品池活动
*/
type AlibabaWdkMarketingItempoolCreateactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolCreateactivityResponse
}

type AlibabaWdkMarketingItempoolCreateactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_createactivity_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
