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
type TaobaoOpenAccountUpdateAPIRequest struct {
    model.Params
    // Open Account
    _paramList   []OpenAccount
}

// 初始化TaobaoOpenAccountUpdateAPIRequest对象
func NewTaobaoOpenAccountUpdateRequest() *TaobaoOpenAccountUpdateAPIRequest{
    return &TaobaoOpenAccountUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.open.account.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// Open Account
func (r *TaobaoOpenAccountUpdateAPIRequest) SetParamList(_paramList []OpenAccount) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r TaobaoOpenAccountUpdateAPIRequest) GetParamList() []OpenAccount {
    return r._paramList
}
