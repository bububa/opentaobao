package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺账户实名校验接口 APIRequest
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

func NewTaobaoSebpIsvUserCheckRequest() *TaobaoSebpIsvUserCheckRequest{
    return &TaobaoSebpIsvUserCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSebpIsvUserCheckRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.user.check"
}

func (r TaobaoSebpIsvUserCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSebpIsvUserCheckRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r TaobaoSebpIsvUserCheckRequest) GetUserName() string {
    return r.userName
}

func (r *TaobaoSebpIsvUserCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoSebpIsvUserCheckRequest) GetName() string {
    return r.name
}

func (r *TaobaoSebpIsvUserCheckRequest) SetIdentity(identity string) error {
    r.identity = identity
    r.Set("identity", identity)
    return nil
}

func (r TaobaoSebpIsvUserCheckRequest) GetIdentity() string {
    return r.identity
}

func (r *TaobaoSebpIsvUserCheckRequest) SetAlipay(alipay string) error {
    r.alipay = alipay
    r.Set("alipay", alipay)
    return nil
}

func (r TaobaoSebpIsvUserCheckRequest) GetAlipay() string {
    return r.alipay
}

