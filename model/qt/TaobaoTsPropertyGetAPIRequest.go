package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotspropertygetAPIRequest 淘宝服务属性查询 API请求
// taobao.ts.property.get
//
// 淘宝服务属性查询
type TaobaotspropertygetAPIRequest struct {
	model.Params
	// 服务收费项code
	_serviceItemCode string
}

// NewTaobaotspropertygetRequest 初始化TaobaotspropertygetAPIRequest对象
func NewTaobaotspropertygetRequest() *TaobaotspropertygetAPIRequest {
	return &TaobaotspropertygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotspropertygetAPIRequest) GetApiMethodName() string {
	return "taobao.ts.property.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotspropertygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotspropertygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceItemCode is ServiceItemCode Setter
// 服务收费项code
func (r *TaobaotspropertygetAPIRequest) SetServiceItemCode(_serviceItemCode string) error {
	r._serviceItemCode = _serviceItemCode
	r.Set("service_item_code", _serviceItemCode)
	return nil
}

// GetServiceItemCode ServiceItemCode Getter
func (r TaobaotspropertygetAPIRequest) GetServiceItemCode() string {
	return r._serviceItemCode
}
