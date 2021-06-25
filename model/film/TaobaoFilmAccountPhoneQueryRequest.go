package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据手机查询匹配账号列表 APIRequest
taobao.film.account.phone.query

根据手机号查询匹配的账号列表
*/
type TaobaoFilmAccountPhoneQueryRequest struct {
    model.Params

    // 11位手机号码
    phone   string 

}

func NewTaobaoFilmAccountPhoneQueryRequest() *TaobaoFilmAccountPhoneQueryRequest{
    return &TaobaoFilmAccountPhoneQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmAccountPhoneQueryRequest) GetApiMethodName() string {
    return "taobao.film.account.phone.query"
}

func (r TaobaoFilmAccountPhoneQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmAccountPhoneQueryRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r TaobaoFilmAccountPhoneQueryRequest) GetPhone() string {
    return r.phone
}

