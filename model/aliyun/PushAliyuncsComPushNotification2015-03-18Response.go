package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送通知 APIResponse
push.aliyuncs.com.pushNotification.2015-03-18

pushNotification
*/
type PushAliyuncsComPushNotification2015-03-18APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"push_aliyuncs_com_pushNotification_2015-03-18_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"