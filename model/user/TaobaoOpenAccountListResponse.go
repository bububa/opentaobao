package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount账号信息查询 APIResponse
taobao.open.account.list

OpenAccount账号信息查询
*/
type TaobaoOpenAccountListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"open_account_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回信息
    
    Datas   []OpenaccountObject `json:"datas,omitempty" xml:"