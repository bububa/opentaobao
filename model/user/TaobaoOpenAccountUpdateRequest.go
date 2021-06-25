package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account数据更新 APIRequest
taobao.open.account.update

Open Account数据更新
*/
type TaobaoOpenAccountUpdateRequest struct {
    model.Params

    // Open Account
    paramList   []OpenAccount 

}

func NewTaobaoOpenAccountUpdateRequest() *TaobaoOpenAccountUpdateRequest{
    return &TaobaoOpenAccountUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountUpdateRequest) GetApiMethodName() string {
    return "taobao.open.account.update"
}

func (r TaobaoOpenAccountUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountUpdateRequest) SetParamList(paramList []OpenAccount) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r TaobaoOpenAccountUpdateRequest) GetParamList() []OpenAccount {
    return r.paramList
}

