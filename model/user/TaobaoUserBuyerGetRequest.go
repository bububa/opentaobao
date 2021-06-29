package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家信息API API请求
taobao.user.buyer.get

查询买家信息API，只能买家类应用调用。
*/
type TaobaoUserBuyerGetRequest struct {
    model.Params
    // 只返回nick, avatar参数
    _fields   string
}

// 初始化TaobaoUserBuyerGetRequest对象
func NewTaobaoUserBuyerGetRequest() *TaobaoUserBuyerGetRequest{
    return &TaobaoUserBuyerGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUserBuyerGetRequest) GetApiMethodName() string {
    return "taobao.user.buyer.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUserBuyerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 只返回nick, avatar参数
func (r *TaobaoUserBuyerGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoUserBuyerGetRequest) GetFields() string {
    return r._fields
}
