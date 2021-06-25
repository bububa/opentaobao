package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取TOP通道解密秘钥 APIRequest
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

func NewTaobaoTopSecretGetRequest() *TaobaoTopSecretGetRequest{
    return &TaobaoTopSecretGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopSecretGetRequest) GetApiMethodName() string {
    return "taobao.top.secret.get"
}

func (r TaobaoTopSecretGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopSecretGetRequest) SetSecretVersion(secretVersion int64) error {
    r.secretVersion = secretVersion
    r.Set("secret_version", secretVersion)
    return nil
}

func (r TaobaoTopSecretGetRequest) GetSecretVersion() int64 {
    return r.secretVersion
}

func (r *TaobaoTopSecretGetRequest) SetRandomNum(randomNum string) error {
    r.randomNum = randomNum
    r.Set("random_num", randomNum)
    return nil
}

func (r TaobaoTopSecretGetRequest) GetRandomNum() string {
    return r.randomNum
}

func (r *TaobaoTopSecretGetRequest) SetCustomerUserId(customerUserId int64) error {
    r.customerUserId = customerUserId
    r.Set("customer_user_id", customerUserId)
    return nil
}

func (r TaobaoTopSecretGetRequest) GetCustomerUserId() int64 {
    return r.customerUserId
}

