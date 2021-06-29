package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺账户实名校验接口 API请求
taobao.sebp.isv.user.check

校验淘小铺账户和身份信息匹配成功
*/
type TaobaoSebpIsvUserCheckRequest struct {
    model.Params
    // 淘宝账号
    _userName   string
    // 姓名
    _name   string
    // 证件号
    _identity   string
    // 支付宝账号
    _alipay   string
}

// 初始化TaobaoSebpIsvUserCheckRequest对象
func NewTaobaoSebpIsvUserCheckRequest() *TaobaoSebpIsvUserCheckRequest{
    return &TaobaoSebpIsvUserCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvUserCheckRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.user.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvUserCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserName Setter
// 淘宝账号
func (r *TaobaoSebpIsvUserCheckRequest) SetUserName(_userName string) error {
    r._userName = _userName
    r.Set("user_name", _userName)
    return nil
}

// UserName Getter
func (r TaobaoSebpIsvUserCheckRequest) GetUserName() string {
    return r._userName
}
// Name Setter
// 姓名
func (r *TaobaoSebpIsvUserCheckRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoSebpIsvUserCheckRequest) GetName() string {
    return r._name
}
// Identity Setter
// 证件号
func (r *TaobaoSebpIsvUserCheckRequest) SetIdentity(_identity string) error {
    r._identity = _identity
    r.Set("identity", _identity)
    return nil
}

// Identity Getter
func (r TaobaoSebpIsvUserCheckRequest) GetIdentity() string {
    return r._identity
}
// Alipay Setter
// 支付宝账号
func (r *TaobaoSebpIsvUserCheckRequest) SetAlipay(_alipay string) error {
    r._alipay = _alipay
    r.Set("alipay", _alipay)
    return nil
}

// Alipay Getter
func (r TaobaoSebpIsvUserCheckRequest) GetAlipay() string {
    return r._alipay
}
