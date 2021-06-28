package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
Open Account导入数据 APIResponse
taobao.open.account.create

Open Account导入数据
*/
type TaobaoOpenAccountCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenAccountCreateResponse `json:"open_account_create_response,omitempty"` 
    TaobaoOpenAccountCreateResponse
}

/* model for simplify = false
type TaobaoOpenAccountCreateResponse struct {

    // 插入数据的Open Account Id的列表
    
    Datas  struct {
        OpenaccountLong  []OpenaccountLong `json:"openaccount_long,omitempty"`
    } `json:"datas,omitempty"`
    

}
*/

type TaobaoOpenAccountCreateResponse struct {

    // 插入数据的Open Account Id的列表
    Datas   []OpenaccountLong `json:"datas,omitempty"`

}
