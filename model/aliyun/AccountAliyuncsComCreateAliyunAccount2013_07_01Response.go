package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建阿里云账号 APIResponse
account.aliyuncs.com.CreateAliyunAccount.2013-07-01

根据给定的阿里云账号，密码以及手机号创建阿里云账号
*/
type AccountAliyuncsComCreateAliyunAccount2013_07_01APIResponse struct {
    model.CommonResponse
    AccountAliyuncsComCreateAliyunAccount2013_07_01Response
}

type AccountAliyuncsComCreateAliyunAccount2013_07_01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_CreateAliyunAccount_2013-07-01_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    ResultData   string `json:"result_data,omitempty" xml:"result_data,omitempty"`

    
}
