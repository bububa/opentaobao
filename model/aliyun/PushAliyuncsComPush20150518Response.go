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
	RequestId     string         `json:"request_id,omitempty" xml:"push_aliyuncs_com_push_20150518_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息ID,用于查询
    
    ResponseParams   string `json:"responseParams,omitempty" xml:"