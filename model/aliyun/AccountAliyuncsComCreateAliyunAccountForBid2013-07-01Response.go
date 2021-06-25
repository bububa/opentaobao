package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
为bid用户创建账号 APIResponse
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记
*/
type AccountAliyuncsComCreateAliyunAccountForBid2013-07-01APIResponse struct {
    model.CommonResponse
    Response *AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Response `json:"account_aliyuncs_com_CreateAliyunAccountForBid_2013-07-01_response,omitempty"`
}

type AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Response struct {

    // 返回结果
    ResultData   string `json:"result_data,omitempty"`

}
