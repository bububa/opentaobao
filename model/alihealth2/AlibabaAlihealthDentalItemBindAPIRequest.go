package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalitembindAPIRequest ISV绑定外部门店id和外部商品id API请求
// alibaba.alihealth.dental.item.bind
//
// ISV绑定外部门店id和外部商品id
type AlibabaalihealthdentalitembindAPIRequest struct {
	model.Params
	// bind_list
	_bindList []StoreItemRelRequest
	// 类型 1 天猫门店 2 支付宝门店
	_type int64
}

// NewAlibabaalihealthdentalitembindRequest 初始化AlibabaalihealthdentalitembindAPIRequest对象
func NewAlibabaalihealthdentalitembindRequest() *AlibabaalihealthdentalitembindAPIRequest {
	return &AlibabaalihealthdentalitembindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalitembindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalitembindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalitembindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindList is BindList Setter
// bind_list
func (r *AlibabaalihealthdentalitembindAPIRequest) SetBindList(_bindList []StoreItemRelRequest) error {
	r._bindList = _bindList
	r.Set("bind_list", _bindList)
	return nil
}

// GetBindList BindList Getter
func (r AlibabaalihealthdentalitembindAPIRequest) GetBindList() []StoreItemRelRequest {
	return r._bindList
}

// SetType is Type Setter
// 类型 1 天猫门店 2 支付宝门店
func (r *AlibabaalihealthdentalitembindAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihealthdentalitembindAPIRequest) GetType() int64 {
	return r._type
}
