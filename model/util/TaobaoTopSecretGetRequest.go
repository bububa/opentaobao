package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取TOP通道解密秘钥 API请求
taobao.top.secret.get

top sdk通过api获取对应解密秘钥
*/
type TaobaoTopSecretGetRequest struct {
    model.Params
    // 秘钥版本号
    _secretVersion   int64
    // 伪随机数
    _randomNum   string
    // 自定义用户id
    _customerUserId   int64
}

// 初始化TaobaoTopSecretGetRequest对象
func NewTaobaoTopSecretGetRequest() *TaobaoTopSecretGetRequest{
    return &TaobaoTopSecretGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretGetRequest) GetApiMethodName() string {
    return "taobao.top.secret.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SecretVersion Setter
// 秘钥版本号
func (r *TaobaoTopSecretGetRequest) SetSecretVersion(_secretVersion int64) error {
    r._secretVersion = _secretVersion
    r.Set("secret_version", _secretVersion)
    return nil
}

// SecretVersion Getter
func (r TaobaoTopSecretGetRequest) GetSecretVersion() int64 {
    return r._secretVersion
}
// RandomNum Setter
// 伪随机数
func (r *TaobaoTopSecretGetRequest) SetRandomNum(_randomNum string) error {
    r._randomNum = _randomNum
    r.Set("random_num", _randomNum)
    return nil
}

// RandomNum Getter
func (r TaobaoTopSecretGetRequest) GetRandomNum() string {
    return r._randomNum
}
// CustomerUserId Setter
// 自定义用户id
func (r *TaobaoTopSecretGetRequest) SetCustomerUserId(_customerUserId int64) error {
    r._customerUserId = _customerUserId
    r.Set("customer_user_id", _customerUserId)
    return nil
}

// CustomerUserId Getter
func (r TaobaoTopSecretGetRequest) GetCustomerUserId() int64 {
    return r._customerUserId
}
