package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount删除数据 APIResponse
taobao.open.account.delete

OpenAccount删除数据
*/
type TaobaoOpenAccountDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenAccountDeleteResponse `json:"taobao_open_account_delete_response,omitempty"`
}

type TaobaoOpenAccountDeleteResponse struct {

    // 删除结果
    Datas   []OpenaccountVoid `json:"datas,omitempty"`

}
