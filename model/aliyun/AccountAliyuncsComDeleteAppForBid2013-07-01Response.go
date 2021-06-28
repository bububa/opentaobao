package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
运营商删除用户的appkey APIResponse
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
*/
type AccountAliyuncsComDeleteAppForBid2013-07-01APIResponse struct {
    model.CommonResponse
    // Response *AccountAliyuncsComDeleteAppForBid2013-07-01Response `json:"account_aliyuncs_com_DeleteAppForBid_2013-07-01_response,omitempty"` 
    AccountAliyuncsComDeleteAppForBid2013-07-01Response
}

/* model for simplify = false
type AccountAliyuncsComDeleteAppForBid2013-07-01Response struct {

    // 用户删除的appkey
    
    AppKey   string `json:"AppKey,omitempty"`
    

    // 调用描述信息
    
    Message   string `json:"Message,omitempty"`
    

    // 结果code
    
    Code   string `json:"Code,omitempty"`
    

    // 入参的requestId
    
    RequestId   string `json:"RequestId,omitempty"`
    

}
*/

type AccountAliyuncsComDeleteAppForBid2013-07-01Response struct {

    // 用户删除的appkey
    AppKey   string `json:"AppKey,omitempty"`

    // 调用描述信息
    Message   string `json:"Message,omitempty"`

    // 结果code
    Code   string `json:"Code,omitempty"`

    // 入参的requestId
    RequestId   string `json:"RequestId,omitempty"`

}
