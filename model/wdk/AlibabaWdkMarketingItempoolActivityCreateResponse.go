package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建活动新接口 APIResponse
alibaba.wdk.marketing.itempool.activity.create

创建活动新接口，支持新工具玩法
*/
type AlibabaWdkMarketingItempoolActivityCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItempoolActivityCreateResponse `json:"alibaba_wdk_marketing_itempool_activity_create_response,omitempty"`
}

type AlibabaWdkMarketingItempoolActivityCreateResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // errorCode
    FailCode   string `json:"fail_code,omitempty"`

    // data
    Data   int64 `json:"data,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
