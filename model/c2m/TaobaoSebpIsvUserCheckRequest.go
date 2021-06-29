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
    userName   string
    // 姓名
    name   string
    // 证件号
    identity   string
    // 支付宝账号
    alipay   string
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
func (r *TaobaoSebpIsvUserCheckRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

// UserName Getter
func (r TaobaoSebpIsvUserCheckRequest) GetUserName() string {
    return r.userName
}
// Name Setter
// 姓名
func (r *TaobaoSebpIsvUserCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoSebpIsvUserCheckRequest) GetName() string {
    return r.name
}
// Identity Setter
// 证件号
func (r *TaobaoSebpIsvUserCheckRequest) SetIdentity(identity string) error {
    r.identity = identity
    r.Set("identity", identity)
    return nil
}

// Identity Getter
func (r TaobaoSebpIsvUserCheckRequest) GetIdentity() string {
    return r.identity
}
// Alipay Setter
// 支付宝账号
func (r *TaobaoSebpIsvUserCheckRequest) SetAlipay(alipay string) error {
    r.alipay = alipay
    r.Set("alipay", alipay)
    return nil
}

// Alipay Getter
func (r TaobaoSebpIsvUserCheckRequest) GetAlipay() string {
    return r.alipay
}
