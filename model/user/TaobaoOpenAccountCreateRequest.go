package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account导入数据 API请求
taobao.open.account.create

Open Account导入数据
*/
type TaobaoOpenAccountCreateRequest struct {
    model.Params
    // Open Account的列表
    _paramList   []OpenAccount
}

// 初始化TaobaoOpenAccountCreateRequest对象
func NewTaobaoOpenAccountCreateRequest() *TaobaoOpenAccountCreateRequest{
    return &TaobaoOpenAccountCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountCreateRequest) GetApiMethodName() string {
    return "taobao.open.account.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// Open Account的列表
func (r *TaobaoOpenAccountCreateRequest) SetParamList(_paramList []OpenAccount) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r TaobaoOpenAccountCreateRequest) GetParamList() []OpenAccount {
    return r._paramList
}
