package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpmbbgetbycodeAPIRequest 根据营销积木块代码获取积木块 API请求
// taobao.ump.mbb.getbycode
//
// 根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
type TaobaoumpmbbgetbycodeAPIRequest struct {
	model.Params
	// 营销积木块code
	_code string
}

// NewTaobaoumpmbbgetbycodeRequest 初始化TaobaoumpmbbgetbycodeAPIRequest对象
func NewTaobaoumpmbbgetbycodeRequest() *TaobaoumpmbbgetbycodeAPIRequest {
	return &TaobaoumpmbbgetbycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpmbbgetbycodeAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbb.getbycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpmbbgetbycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpmbbgetbycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 营销积木块code
func (r *TaobaoumpmbbgetbycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoumpmbbgetbycodeAPIRequest) GetCode() string {
	return r._code
}
