package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
给当前渠道下的用户创建app APIResponse
account.aliyuncs.com.CreateAppForBid.2013-07-01

给自己渠道下的用户创建app
*/
type AccountAliyuncsComCreateAppForBid2013-07-01APIResponse struct {
    model.CommonResponse
    // Response *AccountAliyuncsComCreateAppForBid2013-07-01Response `json:"account_aliyuncs_com_CreateAppForBid_2013-07-01_response,omitempty"` 
    AccountAliyuncsComCreateAppForBid2013-07-01Response
}

/* model for simplify = false
type AccountAliyuncsComCreateAppForBid2013-07-01Response struct {

    // 返回结果
    
    ResultData   string `json:"result_data,omitempty"`
    

}
*/

type AccountAliyuncsComCreateAppForBid2013-07-01Response struct {

    // 返回结果
    ResultData   string `json:"result_data,omitempty"`

}
