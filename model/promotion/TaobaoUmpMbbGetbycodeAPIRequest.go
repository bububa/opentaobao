package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpMbbGetbycodeAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbb.getbycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 营销积木块code
func (r *TaobaoUmpMbbGetbycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoUmpMbbGetbycodeAPIRequest) GetCode() string {
	return r._code
}

var poolTaobaoUmpMbbGetbycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpMbbGetbycodeRequest()
	},
}

// GetTaobaoUmpMbbGetbycodeRequest 从 sync.Pool 获取 TaobaoUmpMbbGetbycodeAPIRequest
func GetTaobaoUmpMbbGetbycodeAPIRequest() *TaobaoUmpMbbGetbycodeAPIRequest {
	return poolTaobaoUmpMbbGetbycodeAPIRequest.Get().(*TaobaoUmpMbbGetbycodeAPIRequest)
}

// ReleaseTaobaoUmpMbbGetbycodeAPIRequest 将 TaobaoUmpMbbGetbycodeAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpMbbGetbycodeAPIRequest(v *TaobaoUmpMbbGetbycodeAPIRequest) {
	v.Reset()
	poolTaobaoUmpMbbGetbycodeAPIRequest.Put(v)
}
