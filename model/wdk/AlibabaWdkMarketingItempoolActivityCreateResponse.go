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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_activity_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"