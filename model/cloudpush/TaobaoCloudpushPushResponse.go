package cloudpush

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户使用云推送高级推送接口 APIResponse
taobao.cloudpush.push

百川用户使用云推送高级推送接口
*/
type TaobaoCloudpushPushAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cloudpush_push_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"