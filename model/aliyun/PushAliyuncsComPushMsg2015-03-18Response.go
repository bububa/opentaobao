package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送 APIResponse
push.aliyuncs.com.pushMsg.2015-03-18

消息推送  ,支持指定用户/账号/广播等模式
*/
type PushAliyuncsComPushMsg2015-03-18APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"push_aliyuncs_com_pushMsg_2015-03-18_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"