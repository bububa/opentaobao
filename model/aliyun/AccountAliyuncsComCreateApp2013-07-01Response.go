package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
给指定用户创建appkey APIResponse
account.aliyuncs.com.CreateApp.2013-07-01

为某个用户创建appkey
*/
type AccountAliyuncsComCreateApp2013-07-01APIResponse struct {
    model.CommonResponse
    Response *AccountAliyuncsComCreateApp2013-07-01Response `json:"account_aliyuncs_com_CreateApp_2013-07-01_response,omitempty"`
}

type AccountAliyuncsComCreateApp2013-07-01Response struct {

    // 返回错误码
    Code   string `json:"code,omitempty"`

    // 返回调用信息
    Message   string `json:"message,omitempty"`

    // 返回结果
    ResultData   string `json:"result_data,omitempty"`

}
