package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建阿里云账号 APIResponse
account.aliyuncs.com.CreateAliyunAccount.2013-07-01

根据给定的阿里云账号，密码以及手机号创建阿里云账号
*/
type AccountAliyuncsComCreateAliyunAccount2013-07-01APIResponse struct {
    model.CommonResponse
    // Response *AccountAliyuncsComCreateAliyunAccount2013-07-01Response `json:"account_aliyuncs_com_CreateAliyunAccount_2013-07-01_response,omitempty"` 
    AccountAliyuncsComCreateAliyunAccount2013-07-01Response
}

/* model for simplify = false
type AccountAliyuncsComCreateAliyunAccount2013-07-01Response struct {

    // 返回结果
    
    ResultData   string `json:"result_data,omitempty"`
    

}
*/

type AccountAliyuncsComCreateAliyunAccount2013-07-01Response struct {

    // 返回结果
    ResultData   string `json:"result_data,omitempty"`

}
