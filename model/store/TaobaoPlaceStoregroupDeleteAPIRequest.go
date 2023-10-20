package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoregroupdeleteAPIRequest 删除门店库 API请求
// taobao.place.storegroup.delete
//
// 删除门店库
type TaobaoplacestoregroupdeleteAPIRequest struct {
	model.Params
	// 库Id
	_id int64
}

// NewTaobaoplacestoregroupdeleteRequest 初始化TaobaoplacestoregroupdeleteAPIRequest对象
func NewTaobaoplacestoregroupdeleteRequest() *TaobaoplacestoregroupdeleteAPIRequest {
	return &TaobaoplacestoregroupdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestoregroupdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestoregroupdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestoregroupdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 库Id
func (r *TaobaoplacestoregroupdeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoplacestoregroupdeleteAPIRequest) GetId() int64 {
	return r._id
}
