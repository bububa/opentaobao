package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account导入数据 APIRequest
taobao.open.account.create

Open Account导入数据
*/
type TaobaoOpenAccountCreateRequest struct {
    model.Params

    // Open Account的列表
    paramList   []OpenAccount 

}

func NewTaobaoOpenAccountCreateRequest() *TaobaoOpenAccountCreateRequest{
    return &TaobaoOpenAccountCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountCreateRequest) GetApiMethodName() string {
    return "taobao.open.account.create"
}

func (r TaobaoOpenAccountCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountCreateRequest) SetParamList(paramList []OpenAccount) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r TaobaoOpenAccountCreateRequest) GetParamList() []OpenAccount {
    return r.paramList
}

