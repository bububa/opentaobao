package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemStoreSchemaUpdateAPIRequest
天猫门店商品编辑 API请求
tmall.item.store.schema.update

天猫门店商品编辑 */
type TmallItemStoreSchemaUpdateAPIRequest struct {
	model.Params
	// 主商品ID
	_mainItemId int64
	// 门店ID
	_storeId int64
	// 商品的schema xml
	_xml string
}

// NewTmallItemStoreSchemaUpdateRequest 初始化TmallItemStoreSchemaUpdateAPIRequest对象
func NewTmallItemStoreSchemaUpdateRequest() *TmallItemStoreSchemaUpdateAPIRequest {
	return &TmallItemStoreSchemaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemStoreSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.store.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemStoreSchemaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MainItemId Setter
// 主商品ID
func (r *TmallItemStoreSchemaUpdateAPIRequest) SetMainItemId(_mainItemId int64) error {
	r._mainItemId = _mainItemId
	r.Set("main_item_id", _mainItemId)
	return nil
}

// Get MainItemId Getter
func (r TmallItemStoreSchemaUpdateAPIRequest) GetMainItemId() int64 {
	return r._mainItemId
}

// Set is StoreId Setter
// 门店ID
func (r *TmallItemStoreSchemaUpdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TmallItemStoreSchemaUpdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is Xml Setter
// 商品的schema xml
func (r *TmallItemStoreSchemaUpdateAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// Get Xml Getter
func (r TmallItemStoreSchemaUpdateAPIRequest) GetXml() string {
	return r._xml
}
