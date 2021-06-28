package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
Open Account数据更新 APIResponse
taobao.open.account.update

Open Account数据更新
*/
type TaobaoOpenAccountUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenAccountUpdateResponse `json:"open_account_update_response,omitempty"` 
    TaobaoOpenAccountUpdateResponse
}

/* model for simplify = false
type TaobaoOpenAccountUpdateResponse struct {

    // update是否成功
    
    Datas  struct {
        OpenaccountVoid  []OpenaccountVoid `json:"openaccount_void,omitempty"`
    } `json:"datas,omitempty"`
    

}
*/

type TaobaoOpenAccountUpdateResponse struct {

    // update是否成功
    Datas   []OpenaccountVoid `json:"datas,omitempty"`

}
