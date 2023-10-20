package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopSecretRegisterAPIRequest) Reset() {
	r.Params.ToZero()
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

var poolTaobaoTopSecretRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopSecretRegisterRequest()
	},
}

// GetTaobaoTopSecretRegisterRequest 从 sync.Pool 获取 TaobaoTopSecretRegisterAPIRequest
func GetTaobaoTopSecretRegisterAPIRequest() *TaobaoTopSecretRegisterAPIRequest {
	return poolTaobaoTopSecretRegisterAPIRequest.Get().(*TaobaoTopSecretRegisterAPIRequest)
}

// ReleaseTaobaoTopSecretRegisterAPIRequest 将 TaobaoTopSecretRegisterAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopSecretRegisterAPIRequest(v *TaobaoTopSecretRegisterAPIRequest) {
	v.Reset()
	poolTaobaoTopSecretRegisterAPIRequest.Put(v)
}
