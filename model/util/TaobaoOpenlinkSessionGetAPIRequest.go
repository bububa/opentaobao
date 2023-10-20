package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenlinkSessionGetAPIRequest 获取授权session信息 API请求
// taobao.openlink.session.get
//
// 帮助第三方isv生成三方session
type TaobaoOpenlinkSessionGetAPIRequest struct {
	model.Params
	// 授权码
	_code string
}

// NewTaobaoOpenlinkSessionGetRequest 初始化TaobaoOpenlinkSessionGetAPIRequest对象
func NewTaobaoOpenlinkSessionGetRequest() *TaobaoOpenlinkSessionGetAPIRequest {
	return &TaobaoOpenlinkSessionGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenlinkSessionGetAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenlinkSessionGetAPIRequest) GetApiMethodName() string {
	return "taobao.openlink.session.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenlinkSessionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenlinkSessionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 授权码
func (r *TaobaoOpenlinkSessionGetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoOpenlinkSessionGetAPIRequest) GetCode() string {
	return r._code
}

var poolTaobaoOpenlinkSessionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenlinkSessionGetRequest()
	},
}

// GetTaobaoOpenlinkSessionGetRequest 从 sync.Pool 获取 TaobaoOpenlinkSessionGetAPIRequest
func GetTaobaoOpenlinkSessionGetAPIRequest() *TaobaoOpenlinkSessionGetAPIRequest {
	return poolTaobaoOpenlinkSessionGetAPIRequest.Get().(*TaobaoOpenlinkSessionGetAPIRequest)
}

// ReleaseTaobaoOpenlinkSessionGetAPIRequest 将 TaobaoOpenlinkSessionGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenlinkSessionGetAPIRequest(v *TaobaoOpenlinkSessionGetAPIRequest) {
	v.Reset()
	poolTaobaoOpenlinkSessionGetAPIRequest.Put(v)
}
