package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemstoreschemaupdateAPIRequest 天猫门店商品编辑 API请求
// tmall.item.store.schema.update
//
// 天猫门店商品编辑
type TmallitemstoreschemaupdateAPIRequest struct {
	model.Params
	// 商品的schema xml
	_xml string
	// 主商品ID
	_mainItemId int64
	// 门店ID
	_storeId int64
}

// NewTmallitemstoreschemaupdateRequest 初始化TmallitemstoreschemaupdateAPIRequest对象
func NewTmallitemstoreschemaupdateRequest() *TmallitemstoreschemaupdateAPIRequest {
	return &TmallitemstoreschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemstoreschemaupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.store.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemstoreschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemstoreschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXml is Xml Setter
// 商品的schema xml
func (r *TmallitemstoreschemaupdateAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r TmallitemstoreschemaupdateAPIRequest) GetXml() string {
	return r._xml
}

// SetMainItemId is MainItemId Setter
// 主商品ID
func (r *TmallitemstoreschemaupdateAPIRequest) SetMainItemId(_mainItemId int64) error {
	r._mainItemId = _mainItemId
	r.Set("main_item_id", _mainItemId)
	return nil
}

// GetMainItemId MainItemId Getter
func (r TmallitemstoreschemaupdateAPIRequest) GetMainItemId() int64 {
	return r._mainItemId
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *TmallitemstoreschemaupdateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TmallitemstoreschemaupdateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
