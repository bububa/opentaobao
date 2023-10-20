package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthreservedentalbindshopanditemAPIRequest 绑定门店信息，商品信息 API请求
// alibaba.alihealth.reserve.dental.bindshopanditem
//
// 绑定门店信息，商品信息
type AlibabaalihealthreservedentalbindshopanditemAPIRequest struct {
	model.Params
	// bind_list
	_bindList []BindDto
}

// NewAlibabaalihealthreservedentalbindshopanditemRequest 初始化AlibabaalihealthreservedentalbindshopanditemAPIRequest对象
func NewAlibabaalihealthreservedentalbindshopanditemRequest() *AlibabaalihealthreservedentalbindshopanditemAPIRequest {
	return &AlibabaalihealthreservedentalbindshopanditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthreservedentalbindshopanditemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.bindshopanditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthreservedentalbindshopanditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthreservedentalbindshopanditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindList is BindList Setter
// bind_list
func (r *AlibabaalihealthreservedentalbindshopanditemAPIRequest) SetBindList(_bindList []BindDto) error {
	r._bindList = _bindList
	r.Set("bind_list", _bindList)
	return nil
}

// GetBindList BindList Getter
func (r AlibabaalihealthreservedentalbindshopanditemAPIRequest) GetBindList() []BindDto {
	return r._bindList
}
