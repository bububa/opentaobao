package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
为bid用户创建账号 APIResponse
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记
*/
type AccountAliyuncsComCreateAliyunAccountForBid2013-07-01APIResponse struct {
    model.CommonResponse
    AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Response
}

type AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_CreateAliyunAccountForBid_2013-07-01_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ResultData   string `json:"result_data,omitempty" xml:"result_data,omitempty"`

    
}
