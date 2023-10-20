package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOauthCodeCreateAPIRequest 淘宝OauthCode颁发 API请求
// taobao.oauth.code.create
//
// 手淘无线开放的oauthCode颁发接口
type TaobaoOauthCodeCreateAPIRequest struct {
	model.Params
	// mock param
	_test int64
}

// NewTaobaoOauthCodeCreateRequest 初始化TaobaoOauthCodeCreateAPIRequest对象
func NewTaobaoOauthCodeCreateRequest() *TaobaoOauthCodeCreateAPIRequest {
	return &TaobaoOauthCodeCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOauthCodeCreateAPIRequest) Reset() {
	r._test = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOauthCodeCreateAPIRequest) GetApiMethodName() string {
	return "taobao.oauth.code.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOauthCodeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOauthCodeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTest is Test Setter
// mock param
func (r *TaobaoOauthCodeCreateAPIRequest) SetTest(_test int64) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// GetTest Test Getter
func (r TaobaoOauthCodeCreateAPIRequest) GetTest() int64 {
	return r._test
}

var poolTaobaoOauthCodeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOauthCodeCreateRequest()
	},
}

// GetTaobaoOauthCodeCreateRequest 从 sync.Pool 获取 TaobaoOauthCodeCreateAPIRequest
func GetTaobaoOauthCodeCreateAPIRequest() *TaobaoOauthCodeCreateAPIRequest {
	return poolTaobaoOauthCodeCreateAPIRequest.Get().(*TaobaoOauthCodeCreateAPIRequest)
}

// ReleaseTaobaoOauthCodeCreateAPIRequest 将 TaobaoOauthCodeCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOauthCodeCreateAPIRequest(v *TaobaoOauthCodeCreateAPIRequest) {
	v.Reset()
	poolTaobaoOauthCodeCreateAPIRequest.Put(v)
}
