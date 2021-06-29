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
    secretVersion   int64
    // 伪随机数
    randomNum   string
    // 自定义用户id
    customerUserId   int64
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
func (r *TaobaoTopSecretGetRequest) SetSecretVersion(secretVersion int64) error {
    r.secretVersion = secretVersion
    r.Set("secret_version", secretVersion)
    return nil
}

// SecretVersion Getter
func (r TaobaoTopSecretGetRequest) GetSecretVersion() int64 {
    return r.secretVersion
}
// RandomNum Setter
// 伪随机数
func (r *TaobaoTopSecretGetRequest) SetRandomNum(randomNum string) error {
    r.randomNum = randomNum
    r.Set("random_num", randomNum)
    return nil
}

// RandomNum Getter
func (r TaobaoTopSecretGetRequest) GetRandomNum() string {
    return r.randomNum
}
// CustomerUserId Setter
// 自定义用户id
func (r *TaobaoTopSecretGetRequest) SetCustomerUserId(customerUserId int64) error {
    r.customerUserId = customerUserId
    r.Set("customer_user_id", customerUserId)
    return nil
}

// CustomerUserId Getter
func (r TaobaoTopSecretGetRequest) GetCustomerUserId() int64 {
    return r.customerUserId
}
