package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemStoreUpdateSchemaGetAPIRequest 天猫门店商品修改规则获取 API请求
// tmall.item.store.update.schema.get
//
// 天猫门店商品修改规则获取
type TmallItemStoreUpdateSchemaGetAPIRequest struct {
	model.Params
	// 主商品ID
	_mainItemId int64
	// 门店ID
	_storeId int64
}

// NewTmallItemStoreUpdateSchemaGetRequest 初始化TmallItemStoreUpdateSchemaGetAPIRequest对象
func NewTmallItemStoreUpdateSchemaGetRequest() *TmallItemStoreUpdateSchemaGetAPIRequest {
	return &TmallItemStoreUpdateSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.store.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainItemId is MainItemId Setter
// 主商品ID
func (r *TmallItemStoreUpdateSchemaGetAPIRequest) SetMainItemId(_mainItemId int64) error {
	r._mainItemId = _mainItemId
	r.Set("main_item_id", _mainItemId)
	return nil
}

// GetMainItemId MainItemId Getter
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetMainItemId() int64 {
	return r._mainItemId
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TmallItemStoreUpdateSchemaGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}
