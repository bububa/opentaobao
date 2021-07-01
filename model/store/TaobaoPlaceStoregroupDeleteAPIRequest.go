package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoregroupDeleteAPIRequest
删除门店库 API请求
taobao.place.storegroup.delete

删除门店库 */
type TaobaoPlaceStoregroupDeleteAPIRequest struct {
	model.Params
	// 库Id
	_id int64
}

// NewTaobaoPlaceStoregroupDeleteRequest 初始化TaobaoPlaceStoregroupDeleteAPIRequest对象
func NewTaobaoPlaceStoregroupDeleteRequest() *TaobaoPlaceStoregroupDeleteAPIRequest {
	return &TaobaoPlaceStoregroupDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.place.storegroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 库Id
func (r *TaobaoPlaceStoregroupDeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TaobaoPlaceStoregroupDeleteAPIRequest) GetId() int64 {
	return r._id
}
