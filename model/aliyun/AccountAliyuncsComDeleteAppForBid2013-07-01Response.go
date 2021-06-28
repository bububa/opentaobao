package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商删除用户的appkey APIResponse
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
*/
type AccountAliyuncsComDeleteAppForBid2013-07-01APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"account_aliyuncs_com_DeleteAppForBid_2013-07-01_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户删除的appkey
    
    AppKey   string `json:"AppKey,omitempty" xml:"