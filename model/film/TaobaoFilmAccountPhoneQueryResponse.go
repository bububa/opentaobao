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
    Response *TaobaoFilmAccountPhoneQueryResponse `json:"taobao_film_account_phone_query_response,omitempty"`
}

type TaobaoFilmAccountPhoneQueryResponse struct {

    // 返回对象
    Result   *ResultListModel `json:"result,omitempty"`

}
