package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建活动新接口 APIResponse
alibaba.wdk.marketing.itempool.activity.create

创建活动新接口，支持新工具玩法
*/
type AlibabaWdkMarketingItempoolActivityCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingItempoolActivityCreateResponse
}

type AlibabaWdkMarketingItempoolActivityCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_activity_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // errorCode
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`

    
    // data
    
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
