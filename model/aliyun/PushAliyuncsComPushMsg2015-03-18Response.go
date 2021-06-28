package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息推送 APIResponse
push.aliyuncs.com.pushMsg.2015-03-18

消息推送  ,支持指定用户/账号/广播等模式
*/
type PushAliyuncsComPushMsg2015-03-18APIResponse struct {
    model.CommonResponse
    // Response *PushAliyuncsComPushMsg2015-03-18Response `json:"push_aliyuncs_com_pushMsg_2015-03-18_response,omitempty"` 
    PushAliyuncsComPushMsg2015-03-18Response
}

/* model for simplify = false
type PushAliyuncsComPushMsg2015-03-18Response struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    
    RequestId   int64 `json:"requestId,omitempty"`
    

    // 消息ID
    
    ResponseParams   int64 `json:"responseParams,omitempty"`
    

}
*/

type PushAliyuncsComPushMsg2015-03-18Response struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    RequestId   int64 `json:"requestId,omitempty"`

    // 消息ID
    ResponseParams   int64 `json:"responseParams,omitempty"`

}
