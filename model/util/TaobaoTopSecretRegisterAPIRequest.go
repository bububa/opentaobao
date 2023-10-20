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
func (r TaobaoTopSecretRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopSecretRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}
