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
type AccountAliyuncsComCreateAppForBid2013-07-01APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"account_aliyuncs_com_CreateAppForBid_2013-07-01_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ResultData   string `json:"result_data,omitempty" xml:"