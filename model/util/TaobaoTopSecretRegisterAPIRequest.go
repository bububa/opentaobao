package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretRegisterAPIRequest 注册加密账号 API请求
// taobao.top.secret.register
//
// 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TaobaoTopSecretRegisterAPIRequest struct {
	model.Params
	// 用户id，保证唯一
	_userId int64
}

// NewTaobaoTopSecretRegisterRequest 初始化TaobaoTopSecretRegisterAPIRequest对象
func NewTaobaoTopSecretRegisterRequest() *TaobaoTopSecretRegisterAPIRequest {
	return &TaobaoTopSecretRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserId is UserId Setter
// 用户id，保证唯一
func (r *TaobaoTopSecretRegisterAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoTopSecretRegisterAPIRequest) GetUserId() int64 {
	return r._userId
}
