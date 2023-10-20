package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorerelatesubdeleteAPIRequest 门店和子门店关系删除 API请求
// taobao.place.storerelatesub.delete
//
// 门店和子门店关系删除
type TaobaoplacestorerelatesubdeleteAPIRequest struct {
	model.Params
	// 子门店id
	_subStoreIds []string
	// 门店Id
	_storeId int64
}

// NewTaobaoplacestorerelatesubdeleteRequest 初始化TaobaoplacestorerelatesubdeleteAPIRequest对象
func NewTaobaoplacestorerelatesubdeleteRequest() *TaobaoplacestorerelatesubdeleteAPIRequest {
	return &TaobaoplacestorerelatesubdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorerelatesubdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.place.storerelatesub.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorerelatesubdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorerelatesubdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubStoreIds is SubStoreIds Setter
// 子门店id
func (r *TaobaoplacestorerelatesubdeleteAPIRequest) SetSubStoreIds(_subStoreIds []string) error {
	r._subStoreIds = _subStoreIds
	r.Set("sub_store_ids", _subStoreIds)
	return nil
}

// GetSubStoreIds SubStoreIds Getter
func (r TaobaoplacestorerelatesubdeleteAPIRequest) GetSubStoreIds() []string {
	return r._subStoreIds
}

// SetStoreId is StoreId Setter
// 门店Id
func (r *TaobaoplacestorerelatesubdeleteAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestorerelatesubdeleteAPIRequest) GetStoreId() int64 {
	return r._storeId
}
