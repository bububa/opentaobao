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
    AccountAliyuncsComCreateApp2013-07-01Response
}

type AccountAliyuncsComCreateApp2013-07-01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_CreateApp_2013-07-01_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回错误码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`

    
    // 返回调用信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回结果
    
    ResultData   string `json:"result_data,omitempty" xml:"result_data,omitempty"`

    
}
