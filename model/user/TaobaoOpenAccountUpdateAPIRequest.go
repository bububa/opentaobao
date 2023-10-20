package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountUpdateAPIRequest Open Account数据更新 API请求
// taobao.open.account.update
//
// Open Account数据更新
type TaobaoOpenAccountUpdateAPIRequest struct {
	model.Params
	// Open Account
	_paramList []OpenAccount
}

// NewTaobaoOpenAccountUpdateRequest 初始化TaobaoOpenAccountUpdateAPIRequest对象
func NewTaobaoOpenAccountUpdateRequest() *TaobaoOpenAccountUpdateAPIRequest {
	return &TaobaoOpenAccountUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenAccountUpdateAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenAccountUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// Open Account
func (r *TaobaoOpenAccountUpdateAPIRequest) SetParamList(_paramList []OpenAccount) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoOpenAccountUpdateAPIRequest) GetParamList() []OpenAccount {
	return r._paramList
}

var poolTaobaoOpenAccountUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenAccountUpdateRequest()
	},
}

// GetTaobaoOpenAccountUpdateRequest 从 sync.Pool 获取 TaobaoOpenAccountUpdateAPIRequest
func GetTaobaoOpenAccountUpdateAPIRequest() *TaobaoOpenAccountUpdateAPIRequest {
	return poolTaobaoOpenAccountUpdateAPIRequest.Get().(*TaobaoOpenAccountUpdateAPIRequest)
}

// ReleaseTaobaoOpenAccountUpdateAPIRequest 将 TaobaoOpenAccountUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenAccountUpdateAPIRequest(v *TaobaoOpenAccountUpdateAPIRequest) {
	v.Reset()
	poolTaobaoOpenAccountUpdateAPIRequest.Put(v)
}
