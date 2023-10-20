package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorerelatesubaddAPIRequest 门店和子门店关系新增 API请求
// taobao.place.storerelatesub.add
//
// 门店和子门店关系新增
type TaobaoplacestorerelatesubaddAPIRequest struct {
	model.Params
	// 子门店Id
	_subStoreIds []string
	// 门店Id
	_storeId int64
}

// NewTaobaoplacestorerelatesubaddRequest 初始化TaobaoplacestorerelatesubaddAPIRequest对象
func NewTaobaoplacestorerelatesubaddRequest() *TaobaoplacestorerelatesubaddAPIRequest {
	return &TaobaoplacestorerelatesubaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorerelatesubaddAPIRequest) GetApiMethodName() string {
	return "taobao.place.storerelatesub.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorerelatesubaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorerelatesubaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubStoreIds is SubStoreIds Setter
// 子门店Id
func (r *TaobaoplacestorerelatesubaddAPIRequest) SetSubStoreIds(_subStoreIds []string) error {
	r._subStoreIds = _subStoreIds
	r.Set("sub_store_ids", _subStoreIds)
	return nil
}

// GetSubStoreIds SubStoreIds Getter
func (r TaobaoplacestorerelatesubaddAPIRequest) GetSubStoreIds() []string {
	return r._subStoreIds
}

// SetStoreId is StoreId Setter
// 门店Id
func (r *TaobaoplacestorerelatesubaddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestorerelatesubaddAPIRequest) GetStoreId() int64 {
	return r._storeId
}
