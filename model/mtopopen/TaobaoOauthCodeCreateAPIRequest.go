package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaooauthcodecreateAPIRequest 淘宝OauthCode颁发 API请求
// taobao.oauth.code.create
//
// 手淘无线开放的oauthCode颁发接口
type TaobaooauthcodecreateAPIRequest struct {
	model.Params
	// mock param
	_test int64
}

// NewTaobaooauthcodecreateRequest 初始化TaobaooauthcodecreateAPIRequest对象
func NewTaobaooauthcodecreateRequest() *TaobaooauthcodecreateAPIRequest {
	return &TaobaooauthcodecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaooauthcodecreateAPIRequest) GetApiMethodName() string {
	return "taobao.oauth.code.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaooauthcodecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaooauthcodecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTest is Test Setter
// mock param
func (r *TaobaooauthcodecreateAPIRequest) SetTest(_test int64) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// GetTest Test Getter
func (r TaobaooauthcodecreateAPIRequest) GetTest() int64 {
	return r._test
}
