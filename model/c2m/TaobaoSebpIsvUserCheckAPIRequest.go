package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSebpIsvUserCheckAPIRequest 淘小铺账户实名校验接口 API请求
// taobao.sebp.isv.user.check
//
// 校验淘小铺账户和身份信息匹配成功
type TaobaoSebpIsvUserCheckAPIRequest struct {
	model.Params
	// 淘宝账号
	_userName string
	// 姓名
	_name string
	// 证件号
	_identity string
	// 支付宝账号
	_alipay string
}

// NewTaobaoSebpIsvUserCheckRequest 初始化TaobaoSebpIsvUserCheckAPIRequest对象
func NewTaobaoSebpIsvUserCheckRequest() *TaobaoSebpIsvUserCheckAPIRequest {
	return &TaobaoSebpIsvUserCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvUserCheckAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.isv.user.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvUserCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserName is UserName Setter
// 淘宝账号
func (r *TaobaoSebpIsvUserCheckAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r TaobaoSebpIsvUserCheckAPIRequest) GetUserName() string {
	return r._userName
}

// SetName is Name Setter
// 姓名
func (r *TaobaoSebpIsvUserCheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoSebpIsvUserCheckAPIRequest) GetName() string {
	return r._name
}

// SetIdentity is Identity Setter
// 证件号
func (r *TaobaoSebpIsvUserCheckAPIRequest) SetIdentity(_identity string) error {
	r._identity = _identity
	r.Set("identity", _identity)
	return nil
}

// GetIdentity Identity Getter
func (r TaobaoSebpIsvUserCheckAPIRequest) GetIdentity() string {
	return r._identity
}

// SetAlipay is Alipay Setter
// 支付宝账号
func (r *TaobaoSebpIsvUserCheckAPIRequest) SetAlipay(_alipay string) error {
	r._alipay = _alipay
	r.Set("alipay", _alipay)
	return nil
}

// GetAlipay Alipay Getter
func (r TaobaoSebpIsvUserCheckAPIRequest) GetAlipay() string {
	return r._alipay
}
