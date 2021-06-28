package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云推送指令API APIResponse
push.aliyuncs.com.push.20150518

阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
*/
type PushAliyuncsComPush20150518APIResponse struct {
    model.CommonResponse
    PushAliyuncsComPush20150518Response
}

type PushAliyuncsComPush20150518Response struct {
    XMLName xml.Name `xml:"push_aliyuncs_com_push_20150518_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息ID,用于查询
    
    ResponseParams   string `json:"responseParams,omitempty" xml:"responseParams,omitempty"`

    
    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`

    
    // 该字段的值由服务端生成,返回给用户方便问题追查与定位。
    
    RequestId   int64 `json:"requestId,omitempty" xml:"requestId,omitempty"`

    
}
