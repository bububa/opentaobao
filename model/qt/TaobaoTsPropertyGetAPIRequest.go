package qt

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTsPropertyGetAPIRequest) GetApiMethodName() string {
	return "taobao.ts.property.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTsPropertyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceItemCode Setter
// 服务收费项code
func (r *TaobaoTsPropertyGetAPIRequest) SetServiceItemCode(_serviceItemCode string) error {
	r._serviceItemCode = _serviceItemCode
	r.Set("service_item_code", _serviceItemCode)
	return nil
}

// Get ServiceItemCode Getter
func (r TaobaoTsPropertyGetAPIRequest) GetServiceItemCode() string {
	return r._serviceItemCode
}
