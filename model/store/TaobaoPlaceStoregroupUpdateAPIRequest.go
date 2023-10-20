package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoregroupupdateAPIRequest 门店库修改基本信息 API请求
// taobao.place.storegroup.update
//
// 门店库修改基本信息
type TaobaoplacestoregroupupdateAPIRequest struct {
	model.Params
	// 库名称
	_name string
	// 库备注
	_desc string
	// 库id
	_id int64
}

// NewTaobaoplacestoregroupupdateRequest 初始化TaobaoplacestoregroupupdateAPIRequest对象
func NewTaobaoplacestoregroupupdateRequest() *TaobaoplacestoregroupupdateAPIRequest {
	return &TaobaoplacestoregroupupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoregroupupdateAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoregroupupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoregroupupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 库名称
func (r *TaobaoplacestoregroupupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoplacestoregroupupdateAPIRequest) GetName() string {
	return r._name
}

// SetDesc is Desc Setter
// 库备注
func (r *TaobaoplacestoregroupupdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoplacestoregroupupdateAPIRequest) GetDesc() string {
	return r._desc
}

// SetId is Id Setter
// 库id
func (r *TaobaoplacestoregroupupdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoplacestoregroupupdateAPIRequest) GetId() int64 {
	return r._id
}
