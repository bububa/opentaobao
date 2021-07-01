package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿房东信息添加 API请求
taobao.xhotel.bnbowner.add

添加和更新民宿房东的信息
*/
type TaobaoXhotelBnbownerAddAPIRequest struct {
    model.Params
    // 添加房东信息的对象
    _addOwnerParam   *AddOwnerParam
}

// 初始化TaobaoXhotelBnbownerAddAPIRequest对象
func NewTaobaoXhotelBnbownerAddRequest() *TaobaoXhotelBnbownerAddAPIRequest{
    return &TaobaoXhotelBnbownerAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbownerAddAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.bnbowner.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbownerAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AddOwnerParam Setter
// 添加房东信息的对象
func (r *TaobaoXhotelBnbownerAddAPIRequest) SetAddOwnerParam(_addOwnerParam *AddOwnerParam) error {
    r._addOwnerParam = _addOwnerParam
    r.Set("add_owner_param", _addOwnerParam)
    return nil
}

// AddOwnerParam Getter
func (r TaobaoXhotelBnbownerAddAPIRequest) GetAddOwnerParam() *AddOwnerParam {
    return r._addOwnerParam
}
