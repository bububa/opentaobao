package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
云推送指令API APIResponse
push.aliyuncs.com.push.20150518

阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
*/
type PushAliyuncsComPush20150518APIResponse struct {
    model.CommonResponse
    // Response *PushAliyuncsComPush20150518Response `json:"push_aliyuncs_com_push_20150518_response,omitempty"` 
    PushAliyuncsComPush20150518Response
}

/* model for simplify = false
type PushAliyuncsComPush20150518Response struct {

    // 消息ID,用于查询
    
    ResponseParams   string `json:"responseParams,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    
    RequestId   int64 `json:"requestId,omitempty"`
    

}
*/

type PushAliyuncsComPush20150518Response struct {

    // 消息ID,用于查询
    ResponseParams   string `json:"responseParams,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    RequestId   int64 `json:"requestId,omitempty"`

}
