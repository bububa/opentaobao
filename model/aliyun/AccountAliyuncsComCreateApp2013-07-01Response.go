package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给指定用户创建appkey APIResponse
account.aliyuncs.com.CreateApp.2013-07-01

为某个用户创建appkey
*/
type AccountAliyuncsComCreateApp2013-07-01APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"account_aliyuncs_com_CreateApp_2013-07-01_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回错误码
    
    Code   string `json:"code,omitempty" xml:"