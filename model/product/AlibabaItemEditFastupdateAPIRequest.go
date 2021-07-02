package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemEditFastupdateAPIRequest 商品编辑增量更新 API请求
// alibaba.item.edit.fastupdate
//
// 商品编辑增量更新;
// <br/>该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
// <br/>新增sku用全量接口alibaba.item.edit.submit，先设置销售属性;
type AlibabaItemEditFastupdateAPIRequest struct {
	model.Params
	// 商品类目ID。若不需要修改商品类目，则不用填写
	_catId int64
	// 产品ID，若不需要修改关联的产品信息，则不需要填写
	_spuId int64
	// 商品ID
	_itemId int64
	// 编辑后的schema信息(增量更新，只填写需要更新的字段)
	_schema string
}

// NewAlibabaItemEditFastupdateRequest 初始化AlibabaItemEditFastupdateAPIRequest对象
func NewAlibabaItemEditFastupdateRequest() *AlibabaItemEditFastupdateAPIRequest {
	return &AlibabaItemEditFastupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemEditFastupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.item.edit.fastupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemEditFastupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCatId is CatId Setter
// 商品类目ID。若不需要修改商品类目，则不用填写
func (r *AlibabaItemEditFastupdateAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetSpuId is SpuId Setter
// 产品ID，若不需要修改关联的产品信息，则不需要填写
func (r *AlibabaItemEditFastupdateAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetSpuId() int64 {
	return r._spuId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemEditFastupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSchema is Schema Setter
// 编辑后的schema信息(增量更新，只填写需要更新的字段)
func (r *AlibabaItemEditFastupdateAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetSchema() string {
	return r._schema
}
