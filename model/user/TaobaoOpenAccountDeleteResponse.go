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
    // Response *TaobaoOpenAccountDeleteResponse `json:"open_account_delete_response,omitempty"` 
    TaobaoOpenAccountDeleteResponse
}

/* model for simplify = false
type TaobaoOpenAccountDeleteResponse struct {

    // 删除结果
    
    Datas  struct {
        OpenaccountVoid  []OpenaccountVoid `json:"openaccount_void,omitempty"`
    } `json:"datas,omitempty"`
    

}
*/

type TaobaoOpenAccountDeleteResponse struct {

    // 删除结果
    Datas   []OpenaccountVoid `json:"datas,omitempty"`

}
