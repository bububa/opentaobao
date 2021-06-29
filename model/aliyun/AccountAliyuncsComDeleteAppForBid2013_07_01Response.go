package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商删除用户的appkey API返回值 
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
*/
type AccountAliyuncsComDeleteAppForBid2013_07_01APIResponse struct {
    model.CommonResponse
    AccountAliyuncsComDeleteAppForBid2013_07_01Response
}

// 运营商删除用户的appkey 成功返回结果
type AccountAliyuncsComDeleteAppForBid2013_07_01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_DeleteAppForBid_2013-07-01_response"`
    // 用户删除的appkey
    AppKey   string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
    // 调用描述信息
    Message   string `json:"Message,omitempty" xml:"Message,omitempty"`
    // 结果code
    Code   string `json:"Code,omitempty" xml:"Code,omitempty"`
    // 入参的requestId
    RequestId   string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}
