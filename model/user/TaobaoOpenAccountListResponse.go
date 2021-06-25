package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount账号信息查询 APIResponse
taobao.open.account.list

OpenAccount账号信息查询
*/
type TaobaoOpenAccountListAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenAccountListResponse `json:"taobao_open_account_list_response,omitempty"`
}

type TaobaoOpenAccountListResponse struct {

    // 返回信息
    Datas   []OpenaccountObject `json:"datas,omitempty"`

}