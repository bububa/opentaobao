package qt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTsPropertyGetAPIRequest 淘宝服务属性查询 API请求
// taobao.ts.property.get
//
// 淘宝服务属性查询
type TaobaoTsPropertyGetAPIRequest struct {
	model.Params
	// 服务收费项code
	_serviceItemCode string
}

// NewTaobaoTsPropertyGetRequest 初始化TaobaoTsPropertyGetAPIRequest对象
func NewTaobaoTsPropertyGetRequest() *TaobaoTsPropertyGetAPIRequest {
	return &TaobaoTsPropertyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTsPropertyGetAPIRequest) Reset() {
	r._serviceItemCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTsPropertyGetAPIRequest) GetApiMethodName() string {
	return "taobao.ts.property.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTsPropertyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTsPropertyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceItemCode is ServiceItemCode Setter
// 服务收费项code
func (r *TaobaoTsPropertyGetAPIRequest) SetServiceItemCode(_serviceItemCode string) error {
	r._serviceItemCode = _serviceItemCode
	r.Set("service_item_code", _serviceItemCode)
	return nil
}

// GetServiceItemCode ServiceItemCode Getter
func (r TaobaoTsPropertyGetAPIRequest) GetServiceItemCode() string {
	return r._serviceItemCode
}

var poolTaobaoTsPropertyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTsPropertyGetRequest()
	},
}

// GetTaobaoTsPropertyGetRequest 从 sync.Pool 获取 TaobaoTsPropertyGetAPIRequest
func GetTaobaoTsPropertyGetAPIRequest() *TaobaoTsPropertyGetAPIRequest {
	return poolTaobaoTsPropertyGetAPIRequest.Get().(*TaobaoTsPropertyGetAPIRequest)
}

// ReleaseTaobaoTsPropertyGetAPIRequest 将 TaobaoTsPropertyGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTsPropertyGetAPIRequest(v *TaobaoTsPropertyGetAPIRequest) {
	v.Reset()
	poolTaobaoTsPropertyGetAPIRequest.Put(v)
}
