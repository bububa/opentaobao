package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account数据更新 API请求
taobao.open.account.update

Open Account数据更新
*/
type TaobaoOpenAccountUpdateRequest struct {
    model.Params
    // Open Account
    _paramList   []OpenAccount
}

// 初始化TaobaoOpenAccountUpdateRequest对象
func NewTaobaoOpenAccountUpdateRequest() *TaobaoOpenAccountUpdateRequest{
    return &TaobaoOpenAccountUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountUpdateRequest) GetApiMethodName() string {
    return "taobao.open.account.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// Open Account
func (r *TaobaoOpenAccountUpdateRequest) SetParamList(_paramList []OpenAccount) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r TaobaoOpenAccountUpdateRequest) GetParamList() []OpenAccount {
    return r._paramList
}
