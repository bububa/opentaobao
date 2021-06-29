package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给当前渠道下的用户创建app APIResponse
account.aliyuncs.com.CreateAppForBid.2013-07-01

给自己渠道下的用户创建app
*/
type AccountAliyuncsComCreateAppForBid2013_07_01APIResponse struct {
    model.CommonResponse
    AccountAliyuncsComCreateAppForBid2013_07_01Response
}

type AccountAliyuncsComCreateAppForBid2013_07_01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_CreateAppForBid_2013-07-01_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    ResultData   string `json:"result_data,omitempty" xml:"result_data,omitempty"`

    
}
