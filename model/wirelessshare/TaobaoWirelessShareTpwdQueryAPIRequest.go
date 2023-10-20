package wirelessshare

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessShareTpwdQueryAPIRequest 查询解析淘口令 API请求
// taobao.wireless.share.tpwd.query
//
// 查询解析淘口令
type TaobaoWirelessShareTpwdQueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// NewTaobaoWirelessShareTpwdQueryRequest 初始化TaobaoWirelessShareTpwdQueryAPIRequest对象
func NewTaobaoWirelessShareTpwdQueryRequest() *TaobaoWirelessShareTpwdQueryAPIRequest {
	return &TaobaoWirelessShareTpwdQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWirelessShareTpwdQueryAPIRequest) Reset() {
	r._passwordContent = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.share.tpwd.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPasswordContent is PasswordContent Setter
// 淘口令
func (r *TaobaoWirelessShareTpwdQueryAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r TaobaoWirelessShareTpwdQueryAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}

var poolTaobaoWirelessShareTpwdQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWirelessShareTpwdQueryRequest()
	},
}

// GetTaobaoWirelessShareTpwdQueryRequest 从 sync.Pool 获取 TaobaoWirelessShareTpwdQueryAPIRequest
func GetTaobaoWirelessShareTpwdQueryAPIRequest() *TaobaoWirelessShareTpwdQueryAPIRequest {
	return poolTaobaoWirelessShareTpwdQueryAPIRequest.Get().(*TaobaoWirelessShareTpwdQueryAPIRequest)
}

// ReleaseTaobaoWirelessShareTpwdQueryAPIRequest 将 TaobaoWirelessShareTpwdQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoWirelessShareTpwdQueryAPIRequest(v *TaobaoWirelessShareTpwdQueryAPIRequest) {
	v.Reset()
	poolTaobaoWirelessShareTpwdQueryAPIRequest.Put(v)
}
