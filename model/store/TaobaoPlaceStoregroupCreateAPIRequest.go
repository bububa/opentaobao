package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoregroupcreateAPIRequest 商户门店库创建接口 API请求
// taobao.place.storegroup.create
//
// 用于商家创建线下门店库
type TaobaoplacestoregroupcreateAPIRequest struct {
	model.Params
	// 库名
	_name string
	// 备注
	_desc string
}

// NewTaobaoplacestoregroupcreateRequest 初始化TaobaoplacestoregroupcreateAPIRequest对象
func NewTaobaoplacestoregroupcreateRequest() *TaobaoplacestoregroupcreateAPIRequest {
	return &TaobaoplacestoregroupcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoregroupcreateAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoregroupcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoregroupcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 库名
func (r *TaobaoplacestoregroupcreateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoplacestoregroupcreateAPIRequest) GetName() string {
	return r._name
}

// SetDesc is Desc Setter
// 备注
func (r *TaobaoplacestoregroupcreateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoplacestoregroupcreateAPIRequest) GetDesc() string {
	return r._desc
}
