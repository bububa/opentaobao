package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretGetAPIRequest 获取TOP通道解密秘钥 API请求
// taobao.top.secret.get
//
// top sdk通过api获取对应解密秘钥
type TaobaoTopSecretGetAPIRequest struct {
	model.Params
	// 伪随机数
	_randomNum string
	// 秘钥版本号
	_secretVersion int64
	// 自定义用户id
	_customerUserId int64
}

// NewTaobaoTopSecretGetRequest 初始化TaobaoTopSecretGetAPIRequest对象
func NewTaobaoTopSecretGetRequest() *TaobaoTopSecretGetAPIRequest {
	return &TaobaoTopSecretGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretGetAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopSecretGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRandomNum is RandomNum Setter
// 伪随机数
func (r *TaobaoTopSecretGetAPIRequest) SetRandomNum(_randomNum string) error {
	r._randomNum = _randomNum
	r.Set("random_num", _randomNum)
	return nil
}

// GetRandomNum RandomNum Getter
func (r TaobaoTopSecretGetAPIRequest) GetRandomNum() string {
	return r._randomNum
}

// SetSecretVersion is SecretVersion Setter
// 秘钥版本号
func (r *TaobaoTopSecretGetAPIRequest) SetSecretVersion(_secretVersion int64) error {
	r._secretVersion = _secretVersion
	r.Set("secret_version", _secretVersion)
	return nil
}

// GetSecretVersion SecretVersion Getter
func (r TaobaoTopSecretGetAPIRequest) GetSecretVersion() int64 {
	return r._secretVersion
}

// SetCustomerUserId is CustomerUserId Setter
// 自定义用户id
func (r *TaobaoTopSecretGetAPIRequest) SetCustomerUserId(_customerUserId int64) error {
	r._customerUserId = _customerUserId
	r.Set("customer_user_id", _customerUserId)
	return nil
}

// GetCustomerUserId CustomerUserId Getter
func (r TaobaoTopSecretGetAPIRequest) GetCustomerUserId() int64 {
	return r._customerUserId
}
