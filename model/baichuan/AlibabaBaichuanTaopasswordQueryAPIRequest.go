package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordQueryAPIRequest 查询解析淘口令 API请求
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
type AlibabaBaichuanTaopasswordQueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// NewAlibabaBaichuanTaopasswordQueryRequest 初始化AlibabaBaichuanTaopasswordQueryAPIRequest对象
func NewAlibabaBaichuanTaopasswordQueryRequest() *AlibabaBaichuanTaopasswordQueryAPIRequest {
	return &AlibabaBaichuanTaopasswordQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBaichuanTaopasswordQueryAPIRequest) Reset() {
	r._passwordContent = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPasswordContent is PasswordContent Setter
// 淘口令
func (r *AlibabaBaichuanTaopasswordQueryAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}

var poolAlibabaBaichuanTaopasswordQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBaichuanTaopasswordQueryRequest()
	},
}

// GetAlibabaBaichuanTaopasswordQueryRequest 从 sync.Pool 获取 AlibabaBaichuanTaopasswordQueryAPIRequest
func GetAlibabaBaichuanTaopasswordQueryAPIRequest() *AlibabaBaichuanTaopasswordQueryAPIRequest {
	return poolAlibabaBaichuanTaopasswordQueryAPIRequest.Get().(*AlibabaBaichuanTaopasswordQueryAPIRequest)
}

// ReleaseAlibabaBaichuanTaopasswordQueryAPIRequest 将 AlibabaBaichuanTaopasswordQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaBaichuanTaopasswordQueryAPIRequest(v *AlibabaBaichuanTaopasswordQueryAPIRequest) {
	v.Reset()
	poolAlibabaBaichuanTaopasswordQueryAPIRequest.Put(v)
}
