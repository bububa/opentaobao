package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据手机查询匹配账号列表 APIResponse
taobao.film.account.phone.query

根据手机号查询匹配的账号列表
*/
type TaobaoFilmAccountPhoneQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilmAccountPhoneQueryResponse `json:"film_account_phone_query_response,omitempty"` 
    TaobaoFilmAccountPhoneQueryResponse
}

/* model for simplify = false
type TaobaoFilmAccountPhoneQueryResponse struct {

    // 返回对象
    
    Result  *struct {
        ResultListModel  *ResultListModel `json:"result_list_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFilmAccountPhoneQueryResponse struct {

    // 返回对象
    Result   *ResultListModel `json:"result,omitempty"`

}
