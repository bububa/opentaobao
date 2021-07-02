package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoregroupUpdateAPIRequest 门店库修改基本信息 API请求
// taobao.place.storegroup.update
//
// 门店库修改基本信息
type TaobaoPlaceStoregroupUpdateAPIRequest struct {
	model.Params
	// 库id
	_id int64
	// 库名称
	_name string
	// 库备注
	_desc string
}

// NewTaobaoPlaceStoregroupUpdateRequest 初始化TaobaoPlaceStoregroupUpdateAPIRequest对象
func NewTaobaoPlaceStoregroupUpdateRequest() *TaobaoPlaceStoregroupUpdateAPIRequest {
	return &TaobaoPlaceStoregroupUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 库id
func (r *TaobaoPlaceStoregroupUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoPlaceStoregroupUpdateAPIRequest) GetId() int64 {
	return r._id
}

// SetName is Name Setter
// 库名称
func (r *TaobaoPlaceStoregroupUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoPlaceStoregroupUpdateAPIRequest) GetName() string {
	return r._name
}

// SetDesc is Desc Setter
// 库备注
func (r *TaobaoPlaceStoregroupUpdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoPlaceStoregroupUpdateAPIRequest) GetDesc() string {
	return r._desc
}
