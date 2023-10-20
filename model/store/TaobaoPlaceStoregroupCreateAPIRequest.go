package store

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoregroupCreateAPIRequest 商户门店库创建接口 API请求
// taobao.place.storegroup.create
//
// 用于商家创建线下门店库
type TaobaoPlaceStoregroupCreateAPIRequest struct {
	model.Params
	// 库名
	_name string
	// 备注
	_desc string
}

// NewTaobaoPlaceStoregroupCreateRequest 初始化TaobaoPlaceStoregroupCreateAPIRequest对象
func NewTaobaoPlaceStoregroupCreateRequest() *TaobaoPlaceStoregroupCreateAPIRequest {
	return &TaobaoPlaceStoregroupCreateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoregroupCreateAPIRequest) Reset() {
	r._name = ""
	r._desc = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 库名
func (r *TaobaoPlaceStoregroupCreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetName() string {
	return r._name
}

// SetDesc is Desc Setter
// 备注
func (r *TaobaoPlaceStoregroupCreateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetDesc() string {
	return r._desc
}

var poolTaobaoPlaceStoregroupCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoregroupCreateRequest()
	},
}

// GetTaobaoPlaceStoregroupCreateRequest 从 sync.Pool 获取 TaobaoPlaceStoregroupCreateAPIRequest
func GetTaobaoPlaceStoregroupCreateAPIRequest() *TaobaoPlaceStoregroupCreateAPIRequest {
	return poolTaobaoPlaceStoregroupCreateAPIRequest.Get().(*TaobaoPlaceStoregroupCreateAPIRequest)
}

// ReleaseTaobaoPlaceStoregroupCreateAPIRequest 将 TaobaoPlaceStoregroupCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoregroupCreateAPIRequest(v *TaobaoPlaceStoregroupCreateAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoregroupCreateAPIRequest.Put(v)
}
