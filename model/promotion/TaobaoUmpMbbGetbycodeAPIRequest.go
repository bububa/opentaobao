package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbGetbycodeAPIRequest 根据营销积木块代码获取积木块 API请求
// taobao.ump.mbb.getbycode
//
// 根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
type TaobaoUmpMbbGetbycodeAPIRequest struct {
	model.Params
	// 营销积木块code
	_code string
}

// NewTaobaoUmpMbbGetbycodeRequest 初始化TaobaoUmpMbbGetbycodeAPIRequest对象
func NewTaobaoUmpMbbGetbycodeRequest() *TaobaoUmpMbbGetbycodeAPIRequest {
	return &TaobaoUmpMbbGetbycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbb.getbycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 营销积木块code
func (r *TaobaoUmpMbbGetbycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetCode() string {
	return r._code
}
