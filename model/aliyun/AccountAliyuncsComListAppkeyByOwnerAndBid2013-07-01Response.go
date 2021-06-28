package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据渠道id和状态列出appkey APIResponse
account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01

根据渠道id和状态列出appkey
*/
type AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01APIResponse struct {
    model.CommonResponse
    // Response *AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Response `json:"account_aliyuncs_com_ListAppkeyByOwnerAndBid_2013-07-01_response,omitempty"` 
    AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Response
}

/* model for simplify = false
type AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Response struct {

    // return result
    
    ResultData   string `json:"result_data,omitempty"`
    

}
*/

type AccountAliyuncsComListAppkeyByOwnerAndBid2013-07-01Response struct {

    // return result
    ResultData   string `json:"result_data,omitempty"`

}
